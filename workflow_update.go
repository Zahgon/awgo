// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package aw

// Updater can check for and download & install newer versions of the workflow.
// There is a concrete implementation and documentation in subpackage update.
type Updater interface {
	UpdateAvailable() bool // Return true if a newer version is available
	CheckDue() bool        // Return true if a check for a newer version is due
	CheckForUpdate() error // Retrieve available releases, e.g. from a URL
	Install() error        // Install the latest version
}

// --------------------------------------------------------------------
// Updating

// setUpdater sets an updater for the workflow.
func (wf *Workflow) setUpdater(u Updater) { _ = "STUB: not implemented"; return }

// UpdateCheckDue returns true if an update is available.
func (wf *Workflow) UpdateCheckDue() bool { _ = "STUB: not implemented"; return false }

// CheckForUpdate retrieves and caches the list of available releases.
func (wf *Workflow) CheckForUpdate() error { _ = "STUB: not implemented"; return nil }

// UpdateAvailable returns true if a newer version is available to install.
func (wf *Workflow) UpdateAvailable() bool { _ = "STUB: not implemented"; return false }

// InstallUpdate downloads and installs the latest version of the workflow.
func (wf *Workflow) InstallUpdate() error { _ = "STUB: not implemented"; return nil }
