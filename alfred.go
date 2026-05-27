// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package aw

// JXA scripts to call Alfred.
const (
	scriptSearch = "Application(%s).search(%s);"
	scriptAction = "Application(%s).action(%s);"
	// support "asType" option added in Alfred 4.5
	scriptActionType = "Application(%s).action(%s, %s);"
	scriptBrowse     = "Application(%s).browse(%s);"
	scriptSetTheme   = "Application(%s).setTheme(%s);"
	scriptTrigger    = "Application(%s).runTrigger(%s, %s);"
	scriptSetConfig  = "Application(%s).setConfiguration(%s, %s);"
	scriptRmConfig   = "Application(%s).removeConfiguration(%s, %s);"
	scriptReload     = "Application(%s).reloadWorkflow(%s);"
)

/*
Alfred wraps Alfred's AppleScript API, allowing you to open Alfred in
various modes or call External Triggers.

	a := NewAlfred()

	// Open Alfred
	if err := a.Search(""); err != nil {
		// handle error
	}

	// Browse /Applications
	if err := a.Browse("/Applications"); err != nil {
		// handle error
	}
*/
type Alfred struct {
	Env
	// For testing. Set to true to save JXA script to lastScript
	// instead of running it.
	noRunScripts bool
	lastScript   string
}

// NewAlfred creates a new Alfred from the environment.
//
// It accepts one optional Env argument. If an Env is passed, Alfred
// is initialised from that instead of the system environment.
func NewAlfred(env ...Env) *Alfred { _ = "STUB: not implemented"; return nil }

// Search runs Alfred with the given query. Use an empty query to just open Alfred.
func (a *Alfred) Search(query string) error { _ = "STUB: not implemented"; return nil }

// Browse tells Alfred to open path in navigation mode.
func (a *Alfred) Browse(path string) error { _ = "STUB: not implemented"; return nil }

// SetTheme tells Alfred to use the specified theme.
func (a *Alfred) SetTheme(name string) error { _ = "STUB: not implemented"; return nil }

// Action tells Alfred to show Universal Actions for value(s). This calls Alfred.ActionAsType
// with an empty type.
func (a *Alfred) Action(value ...string) error { _ = "STUB: not implemented"; return nil }

// ActionAsType tells Alfred to show Universal Actions for value(s). Type typ
// may be one of "file", "url" or "text", or an empty string to tell Alfred
// to guess the type.
//
// Added in Alfred 4.5
func (a *Alfred) ActionAsType(typ string, value ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// RunTrigger runs an External Trigger in the given workflow. Query may be empty.
//
// It accepts one optional bundleID argument, which is the bundle ID of the
// workflow whose trigger should be run.
// If not specified, it defaults to the current workflow's.
func (a *Alfred) RunTrigger(name, query string, bundleID ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// ReloadWorkflow tells Alfred to reload a workflow from disk.
//
// It accepts one optional bundleID argument, which is the bundle ID of the
// workflow to reload. If not specified, it defaults to the current workflow's.
func (a *Alfred) ReloadWorkflow(bundleID ...string) error { _ = "STUB: not implemented"; return nil }

func (a *Alfred) runScript(script string, arg ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Name of JXA Application for running Alfred
func scriptAppName() string {
	_ = "STUB: not implemented"
	// Alfred 3
	return ""
}

// Alfred 4+
