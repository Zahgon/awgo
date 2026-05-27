// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package aw

import "go.deanishe.net/fuzzy"

// Option is a configuration option for Workflow.
// Pass one or more Options to New() or Workflow.Configure().
//
// An Option returns its inverse (i.e. an Option that restores the
// previous value).
//
// You can apply Options at any time, so you can, e.g. suppress UIDs
// if you need to for items to be in a particular order.
type Option func(wf *Workflow) Option

// options combines Options, allowing the application of their inverse
// via a single call to options.apply().
type options []Option

// apply configures Workflow with all options and returns a single Option
// to reverse all changes.
func (opts options) apply(wf *Workflow) Option { _ = "STUB: not implemented"; return *new(Option) }

// HelpURL sets link shown in debugger & log if Run() catches a panic
// ("Get help at http://…").
// Set this to the URL of an issue tracker/forum thread where users can
// ask for help.
func HelpURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// LogPrefix is the printed to debugger at the start of each run.
// Its purpose is to ensure that the first real log message is shown
// on its own line.
// It is only sent to Alfred's debugger, not the log file.
//
// Default: Beer Mug (\U0001F37A)
func LogPrefix(prefix string) Option { _ = "STUB: not implemented"; return *new(Option) }

// MagicPrefix sets the prefix for "magic" commands.
// If a user enters this prefix, AwGo takes control of the workflow and
// shows a list of matching magic commands to the user.
//
// Default: workflow:
func MagicPrefix(prefix string) Option { _ = "STUB: not implemented"; return *new(Option) }

// MaxLogSize sets the size (in bytes) when workflow log is rotated.
// Default: 1 MiB
func MaxLogSize(bytes int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MaxResults is the maximum number of results to send to Alfred.
// 0 means send all results.
// Default: 0
func MaxResults(num int) Option { _ = "STUB: not implemented"; return *new(Option) }

// TextErrors tells Workflow to print errors as text, not JSON.
// Messages are still sent to STDOUT. Set to true if error
// should be captured by Alfred, e.g. if output goes to a Notification.
func TextErrors(on bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// SortOptions sets the fuzzy sorting options for Workflow.Filter().
// See fuzzy and fuzzy.Option for info on (configuring) the sorting
// algorithm.
//
// _examples/fuzzy contains an example workflow using fuzzy sort.
func SortOptions(opts ...fuzzy.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// SessionName changes the name of the variable used to store the session ID.
//
// This is useful if you have multiple Script Filters chained together that
// you don't want to use the same cache.
func SessionName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// SuppressUIDs prevents UIDs from being set on feedback Items.
//
// This turns off Alfred's knowledge, i.e. prevents Alfred from
// applying its own sort, so items will be shown in the
// order you add them.
//
// Useful if you need to force a particular item to the top/bottom.
//
// This setting only applies to Items created *after* it has been
// set.
func SuppressUIDs(on bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// Update sets the updater for the Workflow.
// Panics if a version number isn't set (in Alfred Preferences).
//
// See Updater interface and subpackage update for more documentation.
func Update(updater Updater) Option { _ = "STUB: not implemented"; return *new(Option) }

// AddMagic registers Magic Actions with the Workflow.
// Magic Actions connect special keywords/queries to callback functions.
// See the MagicAction interface for more information.
func AddMagic(actions ...MagicAction) Option { _ = "STUB: not implemented"; return *new(Option) }

// RemoveMagic unregisters Magic Actions with Workflow.
// Magic Actions connect special keywords/queries to callback functions.
// See the MagicAction interface for more information.
func RemoveMagic(actions ...MagicAction) Option { _ = "STUB: not implemented"; return *new(Option) }
