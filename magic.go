// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package aw

/*
MagicAction is a command that is called directly by AwGo (i.e.  your workflow
code is not run) if its keyword is passed in a user query.

To use Magic Actions, it's imperative that your workflow call Workflow.Args().

Calls to Workflow.Args() check the workflow's arguments (os.Args[1:])
for the magic prefix ("workflow:" by default), and hijack control of the
workflow if found.

If an exact keyword match is found (e.g. "workflow:log"), the corresponding
action is executed, and the workflow exits.

If no exact match is found, AwGo runs a Script Filter for the user to
select an action. Hitting TAB or RETURN on an item will run it.

Magic Actions are mainly aimed at making debugging and supporting users easier
(via the built-in actions), but they also provide a simple way to integrate
your own commands that don't need a "real" UI.

For example, setting an Updater on Workflow adds an "update" command that
checks for & installs a new version of the workflow.

# Defaults

There are several built-in magic actions, which are registered by
default:

	<prefix>log         Open workflow's log file in the default app (usually
	                    Console).
	<prefix>data        Open workflow's data directory in the default app
	                    (usually Finder).
	<prefix>cache       Open workflow's data directory in the default app
	                    (usually Finder).
	<prefix>deldata     Delete everything in the workflow's data directory.
	<prefix>delcache    Delete everything in the workflow's cache directory.
	<prefix>reset       Delete everything in the workflow's data and cache directories.
	<prefix>help        Open help URL in default browser.
	                    Only registered if you have set a HelpURL.
	<prefix>update      Check for updates and install a newer version of the
	                    workflow if available.
	                    Only registered if you have configured an Updater.

# Custom Actions

To add custom magicActions, you must register them with your Workflow
*before* you call Workflow.Args()

To do this, configure Workflow with the AddMagic option.
*/
type MagicAction interface {
	// Keyword is what the user must enter to run the action after
	// AwGo has recognised the magic prefix. So if the prefix is
	// "workflow:" (the default), a user must enter the query
	// "workflow:<keyword>" to execute this action.
	Keyword() string

	// Description is shown when a user has entered "magic" mode, but
	// the query does not yet match a keyword.
	Description() string

	// RunText is sent to Alfred and written to the log file &
	// debugger when the action is run.
	RunText() string

	// Run is called when the Magic Action is triggered.
	Run() error
}

// magicActions contains the registered magic actions. See the MagicAction
// interface for full documentation.
type magicActions struct {
	actions map[string]MagicAction
	wf      *Workflow
}

// register adds a MagicAction to the mapping. Previous entries are overwritten.
func (ma *magicActions) register(actions ...MagicAction) { _ = "STUB: not implemented"; return }

// unregister removes a MagicAction from the mapping (based on its keyword).
func (ma *magicActions) unregister(actions ...MagicAction) { _ = "STUB: not implemented"; return }

// args runs a magic action or returns command-line arguments.
// It parses args for magic actions. If it finds one, it takes
// control of your workflow and runs the action. Control is
// not returned to your code.
//
// If no magic actions are found, it returns args.
func (ma *magicActions) args(args []string, prefix string) []string {
	_ = "STUB: not implemented"
	return nil
}

// handleArgs checks args for the magic prefix. Returns args and true if
// it found and handled a magic argument.
func (ma *magicActions) handleArgs(args []string, prefix string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Opens workflow's log file.
type logMA struct {
	wf *Workflow
}

func (a logMA) Keyword() string     { _ = "STUB: not implemented"; return "" }
func (a logMA) Description() string { _ = "STUB: not implemented"; return "" }
func (a logMA) RunText() string     { _ = "STUB: not implemented"; return "" }
func (a logMA) Run() error          { _ = "STUB: not implemented"; return nil }

// Opens workflow's data directory.
type dataMA struct {
	wf *Workflow
}

func (a dataMA) Keyword() string     { _ = "STUB: not implemented"; return "" }
func (a dataMA) Description() string { _ = "STUB: not implemented"; return "" }
func (a dataMA) RunText() string     { _ = "STUB: not implemented"; return "" }
func (a dataMA) Run() error          { _ = "STUB: not implemented"; return nil }

// Opens workflow's cache directory.
type cacheMA struct {
	wf *Workflow
}

func (a cacheMA) Keyword() string     { _ = "STUB: not implemented"; return "" }
func (a cacheMA) Description() string { _ = "STUB: not implemented"; return "" }
func (a cacheMA) RunText() string     { _ = "STUB: not implemented"; return "" }
func (a cacheMA) Run() error          { _ = "STUB: not implemented"; return nil }

// Deletes the contents of the workflow's cache directory.
type clearCacheMA struct {
	wf *Workflow
}

func (a clearCacheMA) Keyword() string     { _ = "STUB: not implemented"; return "" }
func (a clearCacheMA) Description() string { _ = "STUB: not implemented"; return "" }
func (a clearCacheMA) RunText() string     { _ = "STUB: not implemented"; return "" }
func (a clearCacheMA) Run() error          { _ = "STUB: not implemented"; return nil }

// Deletes the contents of the workflow's data directory.
type clearDataMA struct {
	wf *Workflow
}

func (a clearDataMA) Keyword() string     { _ = "STUB: not implemented"; return "" }
func (a clearDataMA) Description() string { _ = "STUB: not implemented"; return "" }
func (a clearDataMA) RunText() string     { _ = "STUB: not implemented"; return "" }
func (a clearDataMA) Run() error          { _ = "STUB: not implemented"; return nil }

// Deletes the contents of the workflow's cache & data directories.
type resetMA struct {
	wf *Workflow
}

func (a resetMA) Keyword() string     { _ = "STUB: not implemented"; return "" }
func (a resetMA) Description() string { _ = "STUB: not implemented"; return "" }
func (a resetMA) RunText() string     { _ = "STUB: not implemented"; return "" }
func (a resetMA) Run() error          { _ = "STUB: not implemented"; return nil }

// Opens URL in default browser.
type helpMA struct {
	wf *Workflow
}

func (a helpMA) Keyword() string     { _ = "STUB: not implemented"; return "" }
func (a helpMA) Description() string { _ = "STUB: not implemented"; return "" }
func (a helpMA) RunText() string     { _ = "STUB: not implemented"; return "" }
func (a helpMA) Run() error          { _ = "STUB: not implemented"; return nil }

// Updates the workflow if a newer release is available.
type updateMA struct {
	updater Updater
}

func (a updateMA) Keyword() string     { _ = "STUB: not implemented"; return "" }
func (a updateMA) Description() string { _ = "STUB: not implemented"; return "" }
func (a updateMA) RunText() string     { _ = "STUB: not implemented"; return "" }
func (a updateMA) Run() error          { _ = "STUB: not implemented"; return nil }
