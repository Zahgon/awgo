// Copyright (c) 2019 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT

package build

import (
	"archive/zip"
)

// Export builds an .alfredworkflow file in directory dest
// from the files in directory src. If src is an empty string,
// "build" is used; if dest is empty, "dist" is used.
//
// The filename of the workflow file is generated automatically from
// the workflow's info.plist and is returned if zipping succeeds.
func Export(src, dest string) (path string, err error) { _ = "STUB: not implemented"; return "", nil }

// recursively copy directory to a temporary directory and return path.
func tempCopy(dir string) (tmpdir string, err error) { _ = "STUB: not implemented"; return "", nil }

// remove values for variables marked as unexported
func removeUnexportedVariables(iplist string) (err error) { _ = "STUB: not implemented"; return nil }

func zipFiles(out *zip.Writer, src string) (err error) { _ = "STUB: not implemented"; return nil }

// Glob is a pattern and (relative) destination directory.
type Glob struct {
	// Pattern is a glob-style pattern to match against filesystem
	Pattern string
	// DestDir is a relative directory within target directory
	// to where files matching Pattern should be linked.
	DestDir string
}

// Globs creates a slice of Globs for patterns.
func Globs(pattern ...string) []Glob { _ = "STUB: not implemented"; return nil }

// SymlinkGlobs symlinks multiple Globs to a directory.
func SymlinkGlobs(destDir string, globs ...Glob) error { _ = "STUB: not implemented"; return nil }

// Symlink creates a symlink to target.
func Symlink(link, target string, relative bool) error { _ = "STUB: not implemented"; return nil }
