// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package util

import (
	"regexp"
)

var (
	rxAlphaNum  = regexp.MustCompile(`[^a-zA-Z0-9.-]+`)
	rxMultiDash = regexp.MustCompile(`-+`)
)

// Slugify makes a string filesystem- and URL-safe.
func Slugify(s string) string { _ = "STUB: not implemented"; return "" }

// fold strips diacritics from string.
func fold(s string) string { _ = "STUB: not implemented"; return "" }

// PrettyPath replaces $HOME with ~ in path
func PrettyPath(path string) string { _ = "STUB: not implemented"; return "" }

// PadLeft pads str to length n by adding pad to its left.
func PadLeft(str, pad string, n int) string { _ = "STUB: not implemented"; return "" }

// PadRight pads str to length n by adding pad to its right.
func PadRight(str, pad string, n int) string { _ = "STUB: not implemented"; return "" }

// Pad pads str to length n by adding pad to both ends.
func Pad(str, pad string, n int) string { _ = "STUB: not implemented"; return "" }
