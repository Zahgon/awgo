// Copyright (c) 2019 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT

package util

import (
	"io"
	"os"
)

// MustExist creates all specified directories and returns the last one.
// Panics if any directory cannot be created.
// All created directories have permission set to 0700.
func MustExist(dirpath ...string) string { _ = "STUB: not implemented"; return "" }

// PathExists checks for the existence of path.
// Panics if an error is encountered.
func PathExists(path string) bool { _ = "STUB: not implemented"; return false }

// ClearDirectory deletes all files within a directory, but not directory itself.
func ClearDirectory(p string) error { _ = "STUB: not implemented"; return nil }

// WriteFile is an atomic version of ioutil.WriteFile.
// It first writes data to a temporary file and renames this to
// filename if the write is successful.
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure tempfile is deleted

func closeOrPanic(c io.Closer) { _ = "STUB: not implemented"; return }
