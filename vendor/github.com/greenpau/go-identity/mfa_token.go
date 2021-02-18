// Copyright 2020 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identity

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/binary"
	"fmt"
	"hash"
	"math"
	"strconv"
	"strings"
	"time"
)

// MfaToken is a puiblic key in a public-private key pair.
type MfaToken struct {
	ID         string    `json:"id,omitempty" xml:"id,omitempty" yaml:"id,omitempty"`
	Type       string    `json:"type,omitempty" xml:"type,omitempty" yaml:"type,omitempty"`
	Algorithm  string    `json:"algorithm,omitempty" xml:"algorithm,omitempty" yaml:"algorithm,omitempty"`
	Comment    string    `json:"comment,omitempty" xml:"comment,omitempty" yaml:"comment,omitempty"`
	Secret     string    `json:"secret,omitempty" xml:"secret,omitempty" yaml:"secret,omitempty"`
	Period     int       `json:"period,omitempty" xml:"period,omitempty" yaml:"period,omitempty"`
	Digits     int       `json:"digits,omitempty" xml:"digits,omitempty" yaml:"digits,omitempty"`
	Expired    bool      `json:"expired,omitempty" xml:"expired,omitempty" yaml:"expired,omitempty"`
	ExpiredAt  time.Time `json:"expired_at,omitempty" xml:"expired_at,omitempty" yaml:"expired_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" yaml:"created_at,omitempty"`
	Disabled   bool      `json:"disabled,omitempty" xml:"disabled,omitempty" yaml:"disabled,omitempty"`
	DisabledAt time.Time `json:"disabled_at,omitempty" xml:"disabled_at,omitempty" yaml:"disabled_at,omitempty"`
}

// NewMfaToken returns an instance of MfaToken.
func NewMfaToken(opts map[string]interface{}) (*MfaToken, error) {
	if opts == nil {
		return nil, fmt.Errorf("no arguments found")
	}
	for _, k := range []string{"secret"} {
		if _, exists := opts[k]; !exists {
			return nil, fmt.Errorf("argument %s not found", k)
		}
	}
	p := &MfaToken{
		ID:        GetRandomString(40),
		Secret:    opts["secret"].(string),
		CreatedAt: time.Now().UTC(),
	}
	if v, exists := opts["comment"]; exists {
		p.Comment = v.(string)
	}

	// Algorithm
	if v, exists := opts["algo"]; exists {
		p.Algorithm = v.(string)
	}
	p.Algorithm = strings.ToLower(p.Algorithm)
	switch p.Algorithm {
	case "":
		p.Algorithm = "sha1"
	case "sha1", "sha256", "sha512":
	default:
		return nil, fmt.Errorf("invalid mfa token algorithm: %s", p.Algorithm)
	}

	// Period
	if v, exists := opts["period"]; exists {
		period := v.(string)
		periodInt, err := strconv.Atoi(period)
		if err != nil {
			return nil, err
		}
		if period != strconv.Itoa(periodInt) {
			return nil, fmt.Errorf("invalid mfa token period value")
		}

		p.Period = periodInt
	}
	if p.Period < 30 || p.Period > 300 {
		return nil, fmt.Errorf("invalid mfa token period value, must be between 30 to 300 seconds, got %d", p.Period)
	}

	// Type
	if v, exists := opts["type"]; exists {
		p.Type = v.(string)
	}
	switch p.Type {
	case "":
		return nil, fmt.Errorf("empty mfa token type")
	case "totp":
	default:
		return nil, fmt.Errorf("invalid mfa token type: %s", p.Type)
	}

	// Digits
	if v, exists := opts["digits"]; exists {
		digits, err := strconv.Atoi(v.(string))
		if err != nil {
			return nil, err
		}
		p.Digits = digits
	} else {
		p.Digits = 6
	}
	if p.Digits < 4 || p.Digits > 8 {
		return nil, fmt.Errorf("mfa digits must be between 4 and 8 digits long")
	}

	// Codes
	var code1, code2 string
	for _, i := range []string{"1", "2"} {
		v, exists := opts["code"+i]
		if !exists {
			return nil, fmt.Errorf("mfa code %s not found", i)
		}
		code := v.(string)
		if code == "" {
			return nil, fmt.Errorf("MFA code %s is empty", i)
		}
		if len(code) < 4 || len(code) > 8 {
			return nil, fmt.Errorf("MFA code %s is not 4-8 characters", i)
		}
		if i == "1" {
			code1 = code
			if err := p.ValidateCodeWithTime(code, time.Now().Add(-time.Second*time.Duration(p.Period)).UTC()); err != nil {
				return nil, fmt.Errorf("MFA code1 %s is invalid", code)
			}
			continue
		} else {
			code2 = code
			if code2 == code1 {
				return nil, fmt.Errorf("MFA code 1 and 2 match")
			}
			if len(code2) != len(code1) {
				return nil, fmt.Errorf("MFA code 1 and 2 have different length")
			}
			if err := p.ValidateCodeWithTime(code, time.Now().UTC()); err != nil {
				return nil, fmt.Errorf("MFA code2 %s is invalid", code)
			}
		}
	}

	return p, nil
}

// Disable disables MfaToken instance.
func (p *MfaToken) Disable() {
	p.Expired = true
	p.ExpiredAt = time.Now().UTC()
	p.Disabled = true
	p.DisabledAt = time.Now().UTC()
}

// ValidateCode validates a passcode
func (p *MfaToken) ValidateCode(code string) error {
	ts := time.Now().UTC()
	return p.ValidateCodeWithTime(code, ts)
}

// ValidateCodeWithTime validates a passcode at a particular time.
func (p *MfaToken) ValidateCodeWithTime(code string, ts time.Time) error {
	code = strings.TrimSpace(code)
	if len(code) != p.Digits {
		return fmt.Errorf("passcode length is invalid")
	}
	tp := uint64(math.Floor(float64(ts.Unix()) / float64(p.Period)))
	tps := []uint64{}
	tps = append(tps, tp)
	tps = append(tps, tp+uint64(1))
	tps = append(tps, tp-uint64(1))
	for _, uts := range tps {
		localCode, err := generateMfaCode(p.Secret, p.Algorithm, p.Digits, uts)
		if err != nil {
			continue
		}
		if subtle.ConstantTimeCompare([]byte(localCode), []byte(code)) == 1 {
			return nil
		}
	}
	return fmt.Errorf("passcode is invalid")
}

func generateMfaCode(secret, algo string, digits int, ts uint64) (string, error) {
	var mac hash.Hash
	secretBytes := []byte(secret)
	switch algo {
	case "sha1":
		mac = hmac.New(sha1.New, secretBytes)
	case "sha256":
		mac = hmac.New(sha256.New, secretBytes)
	case "sha512":
		mac = hmac.New(sha512.New, secretBytes)
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", algo)
	}

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, ts)
	mac.Write(buf)
	sum := mac.Sum(nil)

	off := sum[len(sum)-1] & 0xf
	val := int64(((int(sum[off]) & 0x7f) << 24) |
		((int(sum[off+1] & 0xff)) << 16) |
		((int(sum[off+2] & 0xff)) << 8) |
		(int(sum[off+3]) & 0xff))
	mod := int32(val % int64(math.Pow10(digits)))
	wrap := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(wrap, mod), nil
}
