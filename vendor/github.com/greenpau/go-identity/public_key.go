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
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
	"time"
)

var supportedPublicKeyTypes = map[string]bool{
	"ssh": true,
	"gpg": true,
}

// PublicKey is a puiblic key in a public-private key pair.
type PublicKey struct {
	ID    string `json:"id,omitempty" xml:"id,omitempty" yaml:"id,omitempty"`
	Usage string `json:"usage,omitempty" xml:"usage,omitempty" yaml:"usage,omitempty"`
	// Type is any of the following: dsa, rsa, ecdsa, ed25519
	Type           string    `json:"type,omitempty" xml:"type,omitempty" yaml:"type,omitempty"`
	Fingerprint    string    `json:"fingerprint,omitempty" xml:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	FingerprintMD5 string    `json:"fingerprint_md5,omitempty" xml:"fingerprint_md5,omitempty" yaml:"fingerprint_md5,omitempty"`
	Comment        string    `json:"comment,omitempty" xml:"comment,omitempty" yaml:"comment,omitempty"`
	Payload        string    `json:"payload,omitempty" xml:"payload,omitempty" yaml:"payload,omitempty"`
	OpenSSH        string    `json:"openssh,omitempty" xml:"openssh,omitempty" yaml:"openssh,omitempty"`
	Expired        bool      `json:"expired,omitempty" xml:"expired,omitempty" yaml:"expired,omitempty"`
	ExpiredAt      time.Time `json:"expired_at,omitempty" xml:"expired_at,omitempty" yaml:"expired_at,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" yaml:"created_at,omitempty"`
	Disabled       bool      `json:"disabled,omitempty" xml:"disabled,omitempty" yaml:"disabled,omitempty"`
	DisabledAt     time.Time `json:"disabled_at,omitempty" xml:"disabled_at,omitempty" yaml:"disabled_at,omitempty"`
}

// NewPublicKey returns an instance of PublicKey.
func NewPublicKey(opts map[string]interface{}) (*PublicKey, error) {
	if opts == nil {
		return nil, fmt.Errorf("no arguments found")
	}
	for _, k := range []string{"payload", "usage"} {
		if _, exists := opts[k]; !exists {
			return nil, fmt.Errorf("argument %s not found", k)
		}
	}
	p := &PublicKey{
		ID:        GetRandomString(40),
		Payload:   opts["payload"].(string),
		Usage:     opts["usage"].(string),
		CreatedAt: time.Now().UTC(),
	}
	if err := p.parse(); err != nil {
		return nil, err
	}
	if v, exists := opts["comment"]; exists {
		p.Comment = v.(string)
	}
	return p, nil
}

// Disable disables PublicKey instance.
func (p *PublicKey) Disable() {
	p.Expired = true
	p.ExpiredAt = time.Now().UTC()
	p.Disabled = true
	p.DisabledAt = time.Now().UTC()
}

func (p *PublicKey) parse() error {
	switch p.Usage {
	case "ssh", "gpg":
	default:
		return fmt.Errorf("unsupported usage: %s", p.Usage)
	}
	payloadBytes := bytes.TrimSpace([]byte(p.Payload))
	if strings.Contains(p.Payload, "RSA PUBLIC KEY") {
		// Processing PEM file format
		if p.Usage != "ssh" {
			return fmt.Errorf("usage is ssh while payload is not")
		}
		block, _ := pem.Decode(bytes.TrimSpace([]byte(p.Payload)))
		if block.Type != "RSA PUBLIC KEY" {
			return fmt.Errorf("failed decoding RSA PUBLIC KEY: block type mismatch: %s", block.Type)
		}
		publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return fmt.Errorf("failed x509.ParsePKIXPublicKey: %s", err)
		}
		publicKey, err := ssh.NewPublicKey(publicKeyInterface)
		if err != nil {
			return fmt.Errorf("failed ssh.NewPublicKey: %s", err)
		}
		p.Type = publicKey.Type()
		p.FingerprintMD5 = ssh.FingerprintLegacyMD5(publicKey)
		p.Fingerprint = ssh.FingerprintSHA256(publicKey)
		p.Fingerprint = strings.ReplaceAll(p.Fingerprint, "SHA256:", "")
		p.OpenSSH = string(ssh.MarshalAuthorizedKey(publicKey))
		p.OpenSSH = strings.TrimLeft(p.OpenSSH, p.Type+" ")
		return nil
	}
	if strings.Contains(p.Payload, "PGP PUBLIC KEY") {
		// Processing PEM file format
		if p.Usage != "pgp" {
			return fmt.Errorf("usage is pgp while payload is not ")
		}
		return fmt.Errorf("PGP PUBLIC KEY is unsupported")
	}

	// Attempt parsing as authorized OpenSSH keys
	i := bytes.IndexAny(payloadBytes, " \t")
	if i == -1 {
		i = len(payloadBytes)
	}

	var comment []byte
	payloadBase64 := payloadBytes[:i]
	if len(payloadBase64) < 20 {
		// skip preamble, i.e. ssh-rsa, etc.
		payloadBase64 = bytes.TrimSpace(payloadBytes[i:])
		i = bytes.IndexAny(payloadBase64, " \t")
		if i > 0 {
			comment = bytes.TrimSpace(payloadBase64[i:])
			payloadBase64 = payloadBase64[:i]
		}
		p.OpenSSH = string(payloadBase64)
	}
	k := make([]byte, base64.StdEncoding.DecodedLen(len(payloadBase64)))
	n, err := base64.StdEncoding.Decode(k, payloadBase64)
	if err != nil {
		return fmt.Errorf("failed decoding: %s", err)
	}
	publicKey, err := ssh.ParsePublicKey(k[:n])
	if err != nil {
		return fmt.Errorf("failed parsing public key")
	}
	p.Type = publicKey.Type()
	p.Comment = string(comment)
	p.FingerprintMD5 = ssh.FingerprintLegacyMD5(publicKey)
	p.Fingerprint = ssh.FingerprintSHA256(publicKey)

	// Convert OpenSSH key to RSA PUBLIC KEY
	switch publicKey.Type() {
	case "ssh-rsa":
		publicKeyBytes := publicKey.Marshal()
		parsedPublicKey, err := ssh.ParsePublicKey(publicKeyBytes)
		if err != nil {
			return fmt.Errorf("failed parsing OpenSSH key: %s", err)
		}
		cryptoKey := parsedPublicKey.(ssh.CryptoPublicKey)
		publicCryptoKey := cryptoKey.CryptoPublicKey()
		rsaKey := publicCryptoKey.(*rsa.PublicKey)
		rsaKeyASN1, err := x509.MarshalPKIXPublicKey(rsaKey)
		if err != nil {
			return fmt.Errorf("failed converting a public key to PKIX, ASN.1 DER form: %s", err)
		}
		encodedKey := pem.EncodeToMemory(&pem.Block{
			Type: "RSA PUBLIC KEY",
			//Bytes: x509.MarshalPKCS1PublicKey(rsaKey),
			Bytes: rsaKeyASN1,
		})
		p.Payload = string(encodedKey)
	default:
		return fmt.Errorf("unsupported key type: %s", publicKey.Type())
	}

	return nil
}
