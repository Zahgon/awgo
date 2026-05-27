// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

/*
Workflow workflows retrieves and filters GitHub repos tagged with "alfred-workflow".

It demonstrates the use of the caching and background-process APIs to
provide a responsive workflow by updating the datastore in the background.

It shows results based on cached data (if available), and if the cached data
are out-of-date, starts a background process to refresh the cache.

The Script Filter reloads the results (by setting Workflow.Rerun) until the
cached data have been updated (at which point it's showing the latest data).

This is a very useful idiom for workflows that don't need data that are
absolutely bang up-to-date. The user still gets (potentially out-of-date)
results, preserving the responsiveness of the workflow, and the latest data
are shown as soon as they're available.

This workflow, for example, needs several seconds to retrieve all the
search results from GitHub, as there are multiple pages.

NOTE: As the GitHub search is performed in a background process, the
output is not visible in Alfred's debugger. Enter the query "workflow:log"
to view the log file, where you can see the search progress.
*/
package main

import (
	"flag"
	"time"

	aw "github.com/deanishe/awgo"
	"go.deanishe.net/fuzzy"
)

var (
	cacheName   = "repos.json"      // Filename of cached repo list
	maxResults  = 200               // Number of results sent to Alfred
	maxCacheAge = 180 * time.Minute // How long to cache repo list for

	// Command-line flags
	doDownload bool
	query      string

	// Workflow
	sopts []fuzzy.Option
	wf    *aw.Workflow
)

func init() {
	flag.BoolVar(&doDownload, "download", false, "retrieve list of workflows from GitHub")

	// Set some custom fuzzy search options
	sopts = []fuzzy.Option{
		fuzzy.AdjacencyBonus(10.0),
		fuzzy.LeadingLetterPenalty(-0.1),
		fuzzy.MaxLeadingLetterPenalty(-3.0),
		fuzzy.UnmatchedLetterPenalty(-0.5),
	}
	wf = aw.New(aw.HelpURL("http://www.deanishe.net/"),
		aw.MaxResults(maxResults),
		aw.SortOptions(sopts...))
}

func run() {
	_ = "STUB: not implemented"
	// call to handle any magic actions
	return
}

// Try to load repos

// If the cache has expired, set Rerun (which tells Alfred to re-run the
// workflow), and start the background update process if it isn't already
// running.

// Cache is also "expired" if it doesn't exist. So if there are no
// cached data, show a corresponding message and exit.

// Add results for cached repos

// Filter results against query if user entered one

// Convenience method that shows a warning if there are no results to show.
// Alfred's default behaviour if no results are returned is to show its
// fallback searches, which is also what it does if a workflow errors out.
//
// As such, it's a good idea to display a message in this situation,
// otherwise the user can't tell if the workflow failed or simply found
// no matching results.

// Send results/warning message to Alfred

func main() {
	wf.Run(run)
}
