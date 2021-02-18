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
)

// Name represents human name
type Name struct {
	First     string `json:"first,omitempty" xml:"first,omitempty" yaml:"first,omitempty"`
	Last      string `json:"last,omitempty" xml:"last,omitempty" yaml:"last,omitempty"`
	Middle    string `json:"middle,omitempty" xml:"middle,omitempty" yaml:"middle,omitempty"`
	Preferred string `json:"preferred,omitempty" xml:"preferred,omitempty" yaml:"preferred,omitempty"`
	Nickname  bool   `json:"nickname,omitempty" xml:"nickname,omitempty" yaml:"nickname,omitempty"`
	Confirmed bool   `json:"confirmed,omitempty" xml:"confirmed,omitempty" yaml:"confirmed,omitempty"`
	Primary   bool   `json:"primary,omitempty" xml:"primary,omitempty" yaml:"primary,omitempty"`
	Legal     bool   `json:"legal,omitempty" xml:"legal,omitempty" yaml:"legal,omitempty"`
	Alias     bool   `json:"alias,omitempty" xml:"alias,omitempty" yaml:"alias,omitempty"`
}

// NewName returns an instance of Name.
func NewName() *Name {
	return &Name{}
}

// GetNameClaim returns name field of a claim.
func (n *Name) GetNameClaim() string {
	if n.First != "" && n.Last != "" {
		return fmt.Sprintf("%s, %s", n.Last, n.First)
	}
	return ""
}

// GetFullName returns the primary full name for User.
func (n *Name) GetFullName() string {
	if n.First != "" && n.Last != "" {
		return fmt.Sprintf("%s, %s", n.Last, n.First)
	}
	return ""
}
