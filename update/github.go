// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package update

import (
	"regexp"

	aw "github.com/deanishe/awgo"
)

// matches filename of a compiled Alfred workflow
var rxWorkflowFile = regexp.MustCompile(`\.alfred(\d+)?workflow$`)

// GitHub is a Workflow Option. It sets a Workflow Updater for the specified GitHub repo.
// Repo name should be of the form "username/repo", e.g. "deanishe/alfred-ssh".
func GitHub(repo string) aw.Option { _ = "STUB: not implemented"; return *new(aw.Option) }

// create new Updater option from Source.
func newOption(src Source) aw.Option { _ = "STUB: not implemented"; return *new(aw.Option) }

type source struct {
	URL   string
	dls   []Download
	fetch func(URL string) ([]byte, error)
}

// Downloads implements Source.
func (src *source) Downloads() ([]Download, error) { _ = "STUB: not implemented"; return nil, nil }

// parse GitHub/Gitea releases JSON.
func parseReleases(js []byte) ([]Download, error) { _ = "STUB: not implemented"; return nil, nil }

// Reject releases that are empty or contain multiple files with the same extension.
func isValidRelease(dls []Download) error { _ = "STUB: not implemented"; return nil }
