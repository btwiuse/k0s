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
	"encoding/json"
	"fmt"
	"github.com/greenpau/go-identity/internal/utils"
	"github.com/greenpau/versioned"
	"io/ioutil"
	"strings"
	"sync"
)

var (
	app        *versioned.PackageManager
	appVersion string
	gitBranch  string
	gitCommit  string
	buildUser  string
	buildDate  string
)

func init() {
	app = versioned.NewPackageManager("go-identity")
	app.Description = "go-identity"
	app.Documentation = "https://github.com/greenpau/go-identity"
	app.SetVersion(appVersion, "1.0.19")
	app.SetGitBranch(gitBranch, "master")
	app.SetGitCommit(gitCommit, "v1.0.18-2-gfe16763")
	app.SetBuildUser(buildUser, "")
	app.SetBuildDate(buildDate, "")
}

// Database is user identity database.
type Database struct {
	mu              *sync.RWMutex             `json:"-" xml:"-" yaml:"-"`
	Info            *versioned.PackageManager `json:"-" xml:"-" yaml:"-"`
	Revision        uint64                    `json:"revision,omitempty" xml:"revision,omitempty" yaml:"revision,omitempty"`
	RefEmailAddress map[string]*User          `json:"-" xml:"-" yaml:"-"`
	RefUsername     map[string]*User          `json:"-" xml:"-" yaml:"-"`
	RefID           map[string]*User          `json:"-" xml:"-" yaml:"-"`
	Users           []*User                   `json:"users,omitempty" xml:"users,omitempty" yaml:"users,omitempty"`
}

// NewDatabase return an instance of Database.
func NewDatabase() *Database {
	db := &Database{
		mu:              &sync.RWMutex{},
		Info:            app,
		Revision:        1,
		RefUsername:     make(map[string]*User),
		RefID:           make(map[string]*User),
		RefEmailAddress: make(map[string]*User),
		Users:           []*User{},
	}
	return db
}

// AddUser adds user identity to the database.
func (db *Database) AddUser(user *User) error {
	if err := user.Valid(); err != nil {
		return fmt.Errorf("invalid user, %s", err)
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	for i := 0; i < 10; i++ {
		id := NewID()
		if _, exists := db.RefID[id]; !exists {
			user.ID = id
			break
		}
	}
	username := strings.ToLower(user.Username)
	if _, exists := db.RefUsername[username]; exists {
		return fmt.Errorf("username already exists")
	}

	emailAddresses := []string{}
	if len(user.EmailAddresses) > 0 {
		for _, email := range user.EmailAddresses {
			emailAddress := strings.ToLower(email.Address)
			if _, exists := db.RefEmailAddress[emailAddress]; exists {
				return fmt.Errorf("email address already associated with another user")
			}
			emailAddresses = append(emailAddresses, emailAddress)
		}
	}

	db.RefUsername[username] = user
	db.RefID[user.ID] = user
	for _, emailAddress := range emailAddresses {
		db.RefEmailAddress[emailAddress] = user
	}
	db.Users = append(db.Users, user)
	return nil
}

// AuthenticateUser adds user identity to the database.
func (db *Database) AuthenticateUser(username, password string) (map[string]interface{}, bool, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	username = strings.ToLower(username)
	if _, exists := db.RefUsername[username]; !exists {
		return nil, false, fmt.Errorf("username does not exist")
	}
	user := db.RefUsername[username]
	if user == nil {
		return nil, false, fmt.Errorf("user associated with the username is nil")
	}
	if err := user.VerifyPassword(password); err != nil {
		return nil, false, fmt.Errorf("invalid password")
	}
	userMap := make(map[string]interface{})
	userMap["sub"] = username
	if email := user.GetMailClaim(); email != "" {
		userMap["mail"] = email
	}
	if name := user.GetNameClaim(); name != "" {
		userMap["name"] = name
	}
	if roles := user.GetRolesClaim(); roles != "" {
		userMap["roles"] = roles
	}
	return userMap, true, nil
}

// GetUserByID returns a user by id
func (db *Database) GetUserByID(s string) (*User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	userID := strings.ToLower(s)
	if user, exists := db.RefID[userID]; exists {
		return user, nil
	}
	return nil, fmt.Errorf("not found")
}

// GetUserByUsername returns a user by username
func (db *Database) GetUserByUsername(s string) (*User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	username := strings.ToLower(s)
	if user, exists := db.RefUsername[username]; exists {
		return user, nil
	}
	return nil, fmt.Errorf("not found")
}

// GetUserByEmailAddress returns a liast of users associated with a specific email
// address.
func (db *Database) GetUserByEmailAddress(s string) (*User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	email := strings.ToLower(s)
	if user, exists := db.RefEmailAddress[email]; exists {
		return user, nil
	}
	return nil, fmt.Errorf("not found")
}

// GetUserCount returns user count.
func (db *Database) GetUserCount() int {
	return len(db.Users)
}

// SaveToFile saves database contents to JSON file.
func (db *Database) SaveToFile(fp string) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(fp, []byte(data), 0600); err != nil {
		return fmt.Errorf("failed to write data to %s, error: %s", fp, err)
	}
	return nil
}

// LoadFromFile loads database contents from JSON file.
func (db *Database) LoadFromFile(fp string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	content, err := utils.ReadFileBytes(fp)
	if err != nil {
		return err
	}

	tdb := NewDatabase()
	err = json.Unmarshal(content, tdb)
	if err != nil {
		return err
	}

	if len(tdb.Users) > 0 {
		for _, user := range tdb.Users {
			if err := user.Valid(); err != nil {
				return fmt.Errorf("invalid user %v, %s", user, err)
			}
			username := strings.ToLower(user.Username)
			if _, exists := tdb.RefUsername[username]; exists {
				return fmt.Errorf("duplicate username %s %v", user.Username, user)
			}
			if _, exists := tdb.RefID[user.ID]; exists {
				return fmt.Errorf("duplicate user id: %s %v", user.ID, user)
			}
			tdb.RefUsername[username] = user
			tdb.RefID[user.ID] = user
			if len(user.EmailAddresses) > 0 {
				for _, email := range user.EmailAddresses {
					emailAddress := strings.ToLower(email.Address)
					if _, exists := tdb.RefEmailAddress[emailAddress]; exists {
						return fmt.Errorf("duplicate email address: %s %v", emailAddress, user)
					}
					tdb.RefEmailAddress[emailAddress] = user
				}
			}
		}
	}

	db.Revision = tdb.Revision
	db.RefUsername = tdb.RefUsername
	db.RefID = tdb.RefID
	db.RefEmailAddress = tdb.RefEmailAddress
	db.Users = tdb.Users
	return nil
}

func (db *Database) validateUserIdentity(username, email string) (*User, error) {
	user1, err := db.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	user2, err := db.GetUserByEmailAddress(email)
	if err != nil {
		return nil, err
	}
	if user1.ID != user2.ID {
		return nil, fmt.Errorf("username and email point to a different identity")
	}
	return user1, nil
}

// AddPublicKey adds public key, e.g. GPG or SSH, for a user.
func (db *Database) AddPublicKey(opts map[string]interface{}) error {
	var username, email, payload, keyUsage, comment, fp string
	for _, k := range []string{"username", "email", "key", "key_usage", "file_path"} {
		if _, exists := opts[k]; !exists {
			return fmt.Errorf("Password change required %s input field", k)
		}
		switch k {
		case "username":
			username = opts[k].(string)
		case "email":
			email = opts[k].(string)
		case "key":
			payload = opts[k].(string)
		case "file_path":
			fp = opts[k].(string)
		case "key_usage":
			keyUsage = opts[k].(string)
		}
	}
	if v, exists := opts["comment"]; exists {
		comment = v.(string)
	}

	user, err := db.validateUserIdentity(username, email)
	if err != nil {
		return err
	}

	if err := user.AddPublicKey(keyUsage, payload, comment); err != nil {
		return fmt.Errorf("failed adding public %s key, %s", keyUsage, err)
	}
	if err := db.SaveToFile(fp); err != nil {
		return fmt.Errorf("failed to commit newly added public %s key, %s", keyUsage, err)
	}
	return nil
}

// GetPublicKeys returns a list of public keys associated with a user.
func (db *Database) GetPublicKeys(opts map[string]interface{}) ([]*PublicKey, error) {
	var username, email, keyUsage string
	for _, k := range []string{"username", "email", "key_usage"} {
		if _, exists := opts[k]; !exists {
			return nil, fmt.Errorf("Password change required %s input field", k)
		}
		switch k {
		case "username":
			username = opts[k].(string)
		case "email":
			email = opts[k].(string)
		case "key_usage":
			keyUsage = opts[k].(string)
		}
	}

	user, err := db.validateUserIdentity(username, email)
	if err != nil {
		return nil, err
	}

	keys := []*PublicKey{}
	for _, k := range user.PublicKeys {
		if k.Usage != keyUsage {
			continue
		}
		keys = append(keys, k)
	}
	return keys, nil
}

// DeletePublicKey deletes a public key associated with a user by key id.
func (db *Database) DeletePublicKey(opts map[string]interface{}) error {
	var username, email, keyID, fp string
	for _, k := range []string{"username", "email", "key_id", "file_path"} {
		if _, exists := opts[k]; !exists {
			return fmt.Errorf("Password change required %s input field", k)
		}
		switch k {
		case "username":
			username = opts[k].(string)
		case "email":
			email = opts[k].(string)
		case "key_id":
			keyID = opts[k].(string)
		case "file_path":
			fp = opts[k].(string)
		}
	}

	user, err := db.validateUserIdentity(username, email)
	if err != nil {
		return err
	}

	if err := user.DeletePublicKey(keyID); err != nil {
		return err
	}

	if err := db.SaveToFile(fp); err != nil {
		return fmt.Errorf("failed to commit removal of public key, %s", err)
	}
	return nil
}

// ChangeUserPassword  change user password.
func (db *Database) ChangeUserPassword(opts map[string]interface{}) error {
	var username, email, currentPassword, newPassword, fp string
	for _, k := range []string{"username", "email", "current_password", "new_password", "file_path"} {
		if _, exists := opts[k]; !exists {
			return fmt.Errorf("Password change required %s input field", k)
		}
		switch k {
		case "username":
			username = opts[k].(string)
		case "email":
			email = opts[k].(string)
		case "current_password":
			currentPassword = opts[k].(string)
		case "new_password":
			newPassword = opts[k].(string)
		case "file_path":
			fp = opts[k].(string)
		}
	}

	user, err := db.validateUserIdentity(username, email)
	if err != nil {
		return err
	}

	if err := user.VerifyPassword(currentPassword); err != nil {
		return fmt.Errorf("current password is not valid, %s", err)
	}

	if err := user.AddPassword(newPassword); err != nil {
		return fmt.Errorf("failed setting new password, %s", err)
	}

	if err := db.SaveToFile(fp); err != nil {
		return fmt.Errorf("failed to commit new password, %s", err)
	}

	return nil
}

// AddMfaToken adds MFA token for a user.
func (db *Database) AddMfaToken(opts map[string]interface{}) error {
	var username, email, fp string
	for _, k := range []string{"username", "email", "secret", "file_path"} {
		if _, exists := opts[k]; !exists {
			return fmt.Errorf("Password change required %s input field", k)
		}
		switch k {
		case "username":
			username = opts[k].(string)
		case "email":
			email = opts[k].(string)
		case "file_path":
			fp = opts[k].(string)
		}
	}

	user, err := db.validateUserIdentity(username, email)
	if err != nil {
		return err
	}

	tokenOpts := make(map[string]interface{})
	for k, v := range opts {
		switch k {
		case "username", "email", "file_path":
		default:
			tokenOpts[k] = v
		}
	}

	if err := user.AddMfaToken(tokenOpts); err != nil {
		return fmt.Errorf("failed adding MFA token: %s", err)
	}
	if err := db.SaveToFile(fp); err != nil {
		return fmt.Errorf("failed to commit newly added MFA token: %s", err)
	}
	return nil
}

// GetMfaTokens returns a list of MFA tokens associated with a user.
func (db *Database) GetMfaTokens(opts map[string]interface{}) ([]*MfaToken, error) {
	var username, email string
	for _, k := range []string{"username", "email"} {
		if _, exists := opts[k]; !exists {
			return nil, fmt.Errorf("Password change required %s input field", k)
		}
		switch k {
		case "username":
			username = opts[k].(string)
		case "email":
			email = opts[k].(string)
		}
	}

	user, err := db.validateUserIdentity(username, email)
	if err != nil {
		return nil, err
	}

	tokens := []*MfaToken{}
	for _, token := range user.MfaTokens {
		if token.Disabled {
			continue
		}
		tokens = append(tokens, token)
	}
	return user.MfaTokens, nil
}

// DeleteMfaToken deletes MFA token associated with a user by token id.
func (db *Database) DeleteMfaToken(opts map[string]interface{}) error {
	var username, email, tokenID, fp string
	for _, k := range []string{"username", "email", "token_id", "file_path"} {
		if _, exists := opts[k]; !exists {
			return fmt.Errorf("Password change required %s input field", k)
		}
		switch k {
		case "username":
			username = opts[k].(string)
		case "email":
			email = opts[k].(string)
		case "token_id":
			tokenID = opts[k].(string)
		case "file_path":
			fp = opts[k].(string)
		}
	}

	user, err := db.validateUserIdentity(username, email)
	if err != nil {
		return err
	}

	if err := user.DeleteMfaToken(tokenID); err != nil {
		return err
	}

	if err := db.SaveToFile(fp); err != nil {
		return fmt.Errorf("failed to commit removal of MFA token, %s", err)
	}
	return nil
}
