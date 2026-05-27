// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package aw

import (
	"go.deanishe.net/fuzzy"
)

// --------------------------------------------------------------------
// Feedback

// Rerun tells Alfred to re-run the Script Filter after `secs` seconds.
func (wf *Workflow) Rerun(secs float64) *Workflow { _ = "STUB: not implemented"; return nil }

// Vars returns the workflow variables set on Workflow.Feedback.
// See Feedback.Vars() for more information.
func (wf *Workflow) Vars() map[string]string { _ = "STUB: not implemented"; return nil }

// Var sets the value of workflow variable k on Workflow.Feedback to v.
// See Feedback.Var() for more information.
func (wf *Workflow) Var(k, v string) *Workflow { _ = "STUB: not implemented"; return nil }

// NewItem adds and returns a new feedback Item.
// See Feedback.NewItem() for more information.
func (wf *Workflow) NewItem(title string) *Item { _ = "STUB: not implemented"; return nil }

// NewFileItem adds and returns a new Item pre-populated from path.
// Title and Autocomplete are the base name of the file,
// Subtitle is the path to the file (using "~" for $HOME),
// Valid is true,
// UID and Arg are set to path,
// Type is "file", and
// Icon is the icon of the file at path.
func (wf *Workflow) NewFileItem(path string) *Item { _ = "STUB: not implemented"; return nil }

// NewWarningItem adds and returns a new Feedback Item with the system
// warning icon (exclamation mark on yellow triangle).
func (wf *Workflow) NewWarningItem(title, subtitle string) *Item {
	_ = "STUB: not implemented"
	return nil
}

// IsEmpty returns true if Workflow contains no items.
func (wf *Workflow) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// FatalError displays an error message in Alfred, then calls log.Fatal(),
// terminating the workflow.
func (wf *Workflow) FatalError(err error) { _ = "STUB: not implemented"; return }

// Fatal displays an error message in Alfred, then calls log.Fatal(),
// terminating the workflow.
func (wf *Workflow) Fatal(msg string) { _ = "STUB: not implemented"; return }

// Fatalf displays an error message in Alfred, then calls log.Fatal(),
// terminating the workflow.
func (wf *Workflow) Fatalf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Warn displays a warning message in Alfred immediately. Unlike
// FatalError()/Fatal(), this does not terminate the workflow,
// but you can't send any more results to Alfred.
func (wf *Workflow) Warn(title, subtitle string) *Workflow {
	_ = "STUB: not implemented"
	// Remove any existing items
	return nil
}

// WarnEmpty adds a warning item to feedback if there are no other items.
func (wf *Workflow) WarnEmpty(title, subtitle string) { _ = "STUB: not implemented"; return }

// Filter fuzzy-sorts feedback Items against query and deletes Items that don't match.
func (wf *Workflow) Filter(query string) []*fuzzy.Result { _ = "STUB: not implemented"; return nil }

// SendFeedback sends Script Filter results to Alfred.
//
// Results are output as JSON to STDOUT. As you can output results only once,
// subsequent calls to sending methods are logged and ignored.
//
// The sending methods are:
//
//	SendFeedback()
//	Fatal()
//	Fatalf()
//	FatalError()
//	Warn()
//	WarnEmpty()  // only sends if there are no items
func (wf *Workflow) SendFeedback() *Workflow {
	_ = "STUB: not implemented"
	// Set session ID
	return nil
}

// Truncate Items if maxResults is set
