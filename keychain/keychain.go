// Copyright (c) 2019 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT

// Package keychain implements a simple interface to the macOS Keychain.
// Based on /usr/bin/security.
package keychain

import (
	"errors"
)

// Specific errors returned by the API.
var (
	// Returned by Keychain.Get() and Keychain.Delete() if the specified
	// account doesn't exist.
	ErrNotFound = errors.New("password not found")
	// Used internally. Swallowed by Keychain.Set() if account already exists.
	errDuplicate = errors.New("duplicate password")
)

// Keychain manages macOS Keychain passwords for a specific service.
type Keychain struct {
	service string
}

// New Keychain for specified service.
func New(service string) *Keychain { _ = "STUB: not implemented"; return nil }

// Get password from user's Keychain. Returns ErrNotFound if specified account doesn't exist.
func (kc *Keychain) Get(account string) (password string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Set password in user's Keychain. If the account already exists, it is replaced.
func (kc *Keychain) Set(account, password string) error { _ = "STUB: not implemented"; return nil }

// Delete a password from user's Keychain. Returns ErrNotFound if account doesn't exist.
func (kc *Keychain) Delete(account string) error { _ = "STUB: not implemented"; return nil }

// run executes a Keychain command.
func (kc *Keychain) run(command, account string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Extract password from /usr/bin/security output.
// If the secret is ASCII, output looks like:
//
//	password: "secret"
//
// If the secret is non-ASCII, output looks like:
//
//	password: 0x74C3AB73745F73C3A96372C3A974  "t\303\253st_s\303\251cr\303\251t"
//
// where the first field is 0x + hex-encoded secret.
func parseKeychainPassword(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// remove "password: " prefix

// ASCII password

// hex-encoded password
