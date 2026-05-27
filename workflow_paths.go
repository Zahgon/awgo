// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package aw

// Dir returns the path to the workflow's root directory.
func (wf *Workflow) Dir() string { _ = "STUB: not implemented"; return "" }

// CacheDir returns the path to the workflow's cache directory.
func (wf *Workflow) CacheDir() string { _ = "STUB: not implemented"; return "" }

// OpenCache opens the workflow's cache directory in the default application (usually Finder).
func (wf *Workflow) OpenCache() error { _ = "STUB: not implemented"; return nil }

// ClearCache deletes all files from the workflow's cache directory.
func (wf *Workflow) ClearCache() error { _ = "STUB: not implemented"; return nil }

// DataDir returns the path to the workflow's data directory.
func (wf *Workflow) DataDir() string { _ = "STUB: not implemented"; return "" }

// OpenData opens the workflow's data directory in the default application (usually Finder).
func (wf *Workflow) OpenData() error { _ = "STUB: not implemented"; return nil }

// ClearData deletes all files from the workflow's data directory.
func (wf *Workflow) ClearData() error { _ = "STUB: not implemented"; return nil }

// Reset deletes all workflow data (cache and data directories).
func (wf *Workflow) Reset() error { _ = "STUB: not implemented"; return nil }

// LogFile returns the path to the workflow's log file.
func (wf *Workflow) LogFile() string { _ = "STUB: not implemented"; return "" }

// OpenLog opens the workflow's logfile in the default application (usually Console.app).
func (wf *Workflow) OpenLog() error { _ = "STUB: not implemented"; return nil }

// OpenHelp opens the workflow's help URL (if set) in the default browser.
func (wf *Workflow) OpenHelp() error { _ = "STUB: not implemented"; return nil }

// Try to find workflow root based on presence of info.plist.
func findWorkflowRoot(path string) string { _ = "STUB: not implemented"; return "" }

// directories to look in for info.plist
// avoid duplicates in dirs

// Add path and all its parents to dirs & seen

// Add all paths from path upwards and from
// directory executable is in upwards.

// Return path of first directory that contains an info.plist
