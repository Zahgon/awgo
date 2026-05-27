// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

/*
Workflow update is an example of how to use AwGo's update API.

It demonstrates best practices for handling updates, in particular
loading the list of available releases in a background process and only
showing an "Update is available!" message if the user query is empty.

# Details

Its own version (set in info.plist via Alfred's UI) is 0.2 and it's
pointing to the GitHub repo "deanishe/alfred-ssh" (a completely
different workflow), which is several versions ahead.

The Script Filter code loads the time of the last update check, and if a
check is due, it calls itself with the "-check" flag via AwGo's
background job API.

Because the release updater runs in a background process, its output is
not shown in Alfred's debugger. Enter the query "workflow:log" to
open the log file if you would like to see the output of the updater.

When run with "-check", the program calls Workflow.CheckForUpdate() to
cache the available releases.

After that has completed, subsequent runs of the Script Filter will show
an "Update available!" item (if the query is empty).

Actioning (hitting ↩ or ⌘+1) or completing the item (hitting ⇥)
auto-completes the query to "workflow:update", which is the keyword for
one of AwGo's "magic" actions.

At this point, AwGo will take control of execution, and download and
install the newer version of the workflow (but as it's pointing to a
different workflow's repo, Alfred will install a different workflow
rather than actually updating this one).
*/
package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
	// Command-line arguments
	doCheck bool
	query   string

	// Icon to show if an update is available
	iconAvailable = &aw.Icon{Value: "update-available.png"}
	repo          = "deanishe/alfred-ssh" // GitHub repo
	wf            *aw.Workflow            // Our Workflow struct
	// URL of Alfred metadata.json file to use update.Metadata updater.
	// metadataURL = "https://raw.githubusercontent.com/deanishe/alfred-ssh/master/metadata.json"

	// Fake data for Script Filter
	items = []string{
		"Jacob James",
		"John Cameron",
		"Peter Randolph",
		"Troy Richards",
		"Timothy Anderson",
		"Destiny Singleton",
		"Joshua Hunt",
		"Bridget Rodriguez",
		"Ana Santiago",
		"Luis Davis",
		"Mark Jackson",
		"Joseph Johnson",
		"Ryan Hendricks",
		"Mary Henderson",
		"Timothy Perkins",
		"Mary Henry",
		"Mindy Harrison",
		"Amanda Hawkins",
		"Beverly Brown",
		"Laura Brown",
		"Timothy Patterson",
		"Sandra Murphy",
		"Katherine Reese",
		"Nichole Trevino",
		"David Logan",
		"Allison Thompson",
		"Mark Gibbs",
		"Danielle Willis",
		"Kayla Hill",
		"Kevin Morales",
		"Jeffrey Wheeler",
		"James Bradley",
		"Jeffrey Henry",
		"Nicole Conner",
		"Craig Smith",
		"Steven O'Neill",
		"Cathy Mcknight",
		"Nicolas Waters",
		"Shawn Johnson",
		"Antonio Riley",
	}
)

func init() {
	// Command-line flags
	flag.BoolVar(&doCheck, "check", false, "check for a new version")

	wf = aw.New(update.GitHub(repo))
	// To user metadata.json updater
	// wf = aw.New(update.Metadata(metaDataURL))
}

func run() {
	_ = "STUB: not implemented"
	// call to handle magic actions
	return
}

// Alternate action: Get available releases from remote.

// ----------------------------------------------------------------
// Script Filter
// ----------------------------------------------------------------

// Call self with "check" command if an update is due and a check
// job isn't already running.

// Only show update status if query is empty.

// Turn off UIDs to force this item to the top.
// If UIDs are enabled, Alfred will apply its "knowledge"
// to order the results based on your past usage.

// Notify user of update. As this item is invalid (Valid(false)),
// actioning it expands the query to the Autocomplete value.
// "workflow:update" triggers the updater Magic Action that
// is automatically registered when you configure Workflow with
// an Updater.
//
// If executed, the Magic Action downloads the latest version
// of the workflow and asks Alfred to install it.

// Script Filter results

// Add an extra item to reset update status for demo purposes.
// As with the update notification, this item triggers a Magic
// Action that deletes the cached list of releases.

// Filter results on user query if present

func main() {
	wf.Run(run)
}
