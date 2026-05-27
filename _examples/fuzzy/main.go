// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

/*
Workflow fuzzy is a basic demonstration of AwGo's fuzzy filtering.

It displays and filters the contents of your Downloads directory in Alfred,
and allows you to open files, reveal in Finder or browse in Alfred.
*/
package main

import (
	"os"

	aw "github.com/deanishe/awgo"
)

var (
	// Where we'll look for directories
	startDir = os.ExpandEnv("${HOME}/Downloads")
	// Our Workflow object
	wf *aw.Workflow
)

type file struct {
	Path  string
	IsDir bool
}

// readDir returns the paths to all visible subdirectories of `dirpath`.
func readDir(dir string) (files []file) { _ = "STUB: not implemented"; return nil }

// ignore hidden files

// run executes the Script Filter.
func run() {
	_ = "STUB: not implemented"
	// ----------------------------------------------------------------
	// Handle CLI arguments
	// ----------------------------------------------------------------
	return
}

// You should always use wf.Args() in Script Filters. It contains the
// same as os.Args[1:], but the arguments are first parsed for AwGo's
// magic actions (i.e. "workflow:*" to allow the user to easily open
// the log or data/cache directory).

// ----------------------------------------------------------------
// Load data and create Alfred items
// ----------------------------------------------------------------

// Convenience method. Sets Item title to filename, subtitle
// to shortened path, arg to full path, and icon to file icon.

// Alternate actions

// ----------------------------------------------------------------
// Filter items based on user query
// ----------------------------------------------------------------

// ----------------------------------------------------------------
// Send results to Alfred
// ----------------------------------------------------------------

// Show a warning in Alfred if there are no items

// Send JSON to Alfred. After calling this function, you can't send
// any more results to Alfred.

func main() {
	// Initialise workflow
	wf = aw.New()
	// Call workflow via `Run` wrapper to catch any errors, log them
	// and display an error message in Alfred.
	wf.Run(run)
}
