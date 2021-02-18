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
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
	"time"
)

var supportedPasswordTypes = map[string]bool{
	"bcrypt": true,
}

// Password is a memorized secret, typically a string of characters,
// used to confirm the identity of a user.
type Password struct {
	Purpose    string    `json:"purpose,omitempty" xml:"purpose,omitempty" yaml:"purpose,omitempty"`
	Type       string    `json:"type,omitempty" xml:"type,omitempty" yaml:"type,omitempty"`
	Hash       string    `json:"hash,omitempty" xml:"hash,omitempty" yaml:"hash,omitempty"`
	Cost       int       `json:"cost,omitempty" xml:"cost,omitempty" yaml:"cost,omitempty"`
	Expired    bool      `json:"expired,omitempty" xml:"expired,omitempty" yaml:"expired,omitempty"`
	ExpiredAt  time.Time `json:"expired_at,omitempty" xml:"expired_at,omitempty" yaml:"expired_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" yaml:"created_at,omitempty"`
	Disabled   bool      `json:"disabled,omitempty" xml:"disabled,omitempty" yaml:"disabled,omitempty"`
	DisabledAt time.Time `json:"disabled_at,omitempty" xml:"disabled_at,omitempty" yaml:"disabled_at,omitempty"`
}

// NewPassword returns an instance of Password.
func NewPassword(s string) (*Password, error) {
	p := &Password{
		Purpose:   "generic",
		CreatedAt: time.Now().UTC(),
	}
	if err := p.hashPassword(s); err != nil {
		return nil, err
	}
	return p, nil
}

// Disable disables Password instance.
func (p *Password) Disable() {
	p.Expired = true
	p.ExpiredAt = time.Now().UTC()
	p.Disabled = true
	p.DisabledAt = time.Now().UTC()
}

// HashPassword hashes plain text password. The default hashing method
// is bctypt with cost 10.
func (p *Password) hashPassword(s string) error {
	var password string
	if s == "" {
		return fmt.Errorf("password is empty")
	}
	parts := strings.Split(s, ":")
	if len(parts) == 1 {
		p.Type = "bcrypt"
		password = s
	}
	if len(parts) > 1 {
		p.Type = parts[0]
		password = parts[1]
	}

	if p.Type == "bcrypt" {
		if len(parts) > 2 {
			cost, err := strconv.Atoi(parts[2])
			if err != nil {
				return fmt.Errorf("bcrypt error: failed parsing cost %s", parts[2])
			}
			p.Cost = cost
		}
		if p.Cost == 0 {
			p.Cost = 10
		}
		ph, err := bcrypt.GenerateFromPassword([]byte(password), p.Cost)
		if err != nil {
			return fmt.Errorf("failed hashing password")
		}
		p.Hash = string(ph)
		return nil
	}
	return fmt.Errorf("failed to hash a password, no hashing method found")
}

// Match returns true when the provided password matches the user.
func (p *Password) Match(s string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p.Hash), []byte(s)); err == nil {
		return true
	}
	return false
}
