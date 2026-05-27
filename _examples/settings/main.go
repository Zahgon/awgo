// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

/*
Workflow settings demonstrates binding a struct to Alfred's settings.

The workflow's settings are stored in info.plist/the workflow's
configuration sheet in Alfred Preferences.

These are imported into the Server struct using Config.To().

The Script Filter displays these settings, and you can select one
to change its value.

If you enter a new value, this is saved to info.plist/the configuration
sheet via Config.Set(), and the workflow is run again by calling
its "settings" External Trigger via Alfred.RunTrigger().
*/
package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
)

// Server contains the configuration loaded from the workflow's settings
// in the configuration sheet.
type Server struct {
	Hostname   string
	PortNumber int `env:"PORT"`
	Username   string
	APIKey     string
}

var (
	srv *Server
	wf  *aw.Workflow
	// Command-line arguments
	setKey, getKey string
)

func init() {
	wf = aw.New()
	flag.StringVar(&setKey, "set", "", "save a value")
	flag.StringVar(&getKey, "get", "", "enter a new value")
}

// save setting to info.plist via Alfred's AppleScript API
func runSet(key, value string) { _ = "STUB: not implemented"; return }

// get new value for setting from user via Script Filter
func runGet(key, value string) { _ = "STUB: not implemented"; return }

func run() {
	_ = "STUB: not implemented"
	// call to handle magic actions
	return
}

// ----------------------------------------------------------------
// Load configuration

// Default configuration

// Update config from environment variables

// ----------------------------------------------------------------
// Parse command-line flags and decide what to do

// ----------------------------------------------------------------
// Show available settings.

func main() {
	wf.Run(run)
}
