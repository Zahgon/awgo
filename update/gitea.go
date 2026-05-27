// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package update

import (
	aw "github.com/deanishe/awgo"
)

// Gitea is a Workflow Option. It sets a Workflow Updater for the specified Gitea repo.
// Repo name should be the URL of the repo, e.g. "git.deanishe.net/deanishe/alfred-ssh".
func Gitea(repo string) aw.Option { _ = "STUB: not implemented"; return *new(aw.Option) }

func giteaURL(repo string) string { _ = "STUB: not implemented"; return "" }

// If no scheme is specified, assume HTTPS and re-parse URL.
// This is necessary because URL.Host isn't present on URLs
// without a scheme (hostname is added to path)
