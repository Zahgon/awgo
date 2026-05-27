// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package update

// SemVers implements sort.Interface for SemVer.
type SemVers []SemVer

// Len implements sort.Interface
func (vs SemVers) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface
	return 0
}

func (vs SemVers) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface
func (vs SemVers) Swap(i, j int) { _ = "STUB: not implemented"; return }

// SortSemVer sorts a slice of SemVer structs.
func SortSemVer(versions []SemVer) { _ = "STUB: not implemented"; return }

// SemVer is a (mostly) semantic version number.
//
// Unlike the semver standard:
//   - Minor and patch versions are not required, e.g. "v1" and "v1.0" are valid.
//   - Version string may be prefixed with "v", e.g. "v1" or "v3.0.1-beta".
//     The "v" prefix is stripped, so "v1" == "1.0.0".
//   - Dots and integers are ignored in pre-release identifiers: they are
//     compared purely alphanumerically, e.g. "v1-beta.11" < "v1-beta.2".
//     Use "v1-beta.02" instead.
type SemVer struct {
	Major      uint64 // Increment for breaking changes.
	Minor      uint64 // Increment for added/deprecated functionality.
	Patch      uint64 // Increment for bugfixes.
	Build      string // Build metadata (ignored in comparisons)
	Prerelease string // Pre-release version (treated as string)
}

// NewSemVer creates a new SemVer. An error is returned if the version
// string is not valid. See the SemVer struct documentation for deviations
// from the semver standard.
func NewSemVer(s string) (SemVer, error) { _ = "STUB: not implemented"; return *new(SemVer), nil }

// Remove "v" prefix and extend short versions to full length.

// Extract build and pre tags

// Pad version

// Major

// Minor

// Patch

// String returns a canonical semver string
func (v SemVer) String() string { _ = "STUB: not implemented"; return "" }

// Compare compares two Versions. Returns:
//
//	-1 if v < v2
//	 0 if v == v2
//	 1 if v > v2
func (v SemVer) Compare(v2 SemVer) int { _ = "STUB: not implemented"; return 0 }

// Check if one version is prerelease and the other isn't

// Semver ignores build info

// Eq checks if v == v2
func (v SemVer) Eq(v2 SemVer) bool { _ = "STUB: not implemented"; return false }

// Ne checks if v != v2
func (v SemVer) Ne(v2 SemVer) bool {
	_ = "STUB: not implemented"

	// Gt checks if v > v2
	return false
}

func (v SemVer) Gt(v2 SemVer) bool { _ = "STUB: not implemented"; return false }

// Gte checks if v >= v2
func (v SemVer) Gte(v2 SemVer) bool { _ = "STUB: not implemented"; return false }

// Lt checks if v < v2
func (v SemVer) Lt(v2 SemVer) bool { _ = "STUB: not implemented"; return false }

// Lte checks if v <= v2
func (v SemVer) Lte(v2 SemVer) bool { _ = "STUB: not implemented"; return false }

// IsZero returns true if SemVer has no value.
func (v SemVer) IsZero() bool { _ = "STUB: not implemented"; return false }

func hasLeadingZeroes(s string) bool { _ = "STUB: not implemented"; return false }
