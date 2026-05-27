// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package update

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/deanishe/awgo/util"
)

var (
	// UpdateInterval is how often to check for updates.
	UpdateInterval = 24 * time.Hour
	// HTTPTimeout is the timeout for establishing an HTTP(S) connection.
	HTTPTimeout = 60 * time.Second

	// HTTP client used to talk to APIs
	client *http.Client
)

// Mockable functions
var (
	// Run command
	runCommand = func(name string, arg ...string) error {
		return exec.Command(name, arg...).Run()
	}
	// save a URL to a filepath.
	download = func(URL, path string) error {
		res, err := openURL(URL)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		util.MustExist(filepath.Dir(path))
		out, err := os.Create(path)
		if err != nil {
			return err
		}
		defer out.Close()
		n, err := io.Copy(out, res.Body)
		if err != nil {
			return err
		}
		log.Printf("wrote %q (%d bytes)", util.PrettyPath(path), n)
		return nil
	}
)

// Source provides workflow files that can be downloaded.
// This is what concrete updaters (e.g. GitHub, Gitea) should implement.
// Source is called by the Updater after every updater interval.
type Source interface {
	// Downloads returns all available workflow files.
	Downloads() ([]Download, error)
}

// byVersion sorts downloads by version.
type byVersion []Download

// Len implements sort.Interface.
func (s byVersion) Len() int      { _ = "STUB: not implemented"; return 0 }
func (s byVersion) Swap(i, j int) { _ = "STUB: not implemented"; return }
func (s byVersion) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// Compare workflow versions first, compatible Alfred version second.
	return false
}

// Download is an Alfred workflow available for download & installation.
// It is the primary update data structure, returned by all Sources.
type Download struct {
	URL string // Where the workflow file can be downloaded from
	// Filename for downloaded file.
	// Must have extension .alfredworkflow or .alfredXworkflow where X is a number,
	// otherwise the Download will be ignored.
	Filename   string
	Version    SemVer // Semantic version no.
	Prerelease bool   // Whether this version is a pre-release
}

// AlfredVersion returns minimum compatible version of Alfred based on file extension.
// For example, Workflow.alfred4workflow has version 4, while
// Workflow.alfred3workflow has version 3.
// The standard .alfredworkflow extension returns a zero version.
func (dl Download) AlfredVersion() SemVer { _ = "STUB: not implemented"; return *new(SemVer) }

// Updater checks for newer version of the workflow. Available versions are
// provided by a Source, such as the built-in GitHub source, which
// reads the releases in a GitHub repo. It is a concrete implementation
// of aw.Updater.
//
// CheckForUpdate() retrieves the list of available downloads from the
// source and caches them. UpdateAvailable() reads the cache and returns
// true if there is a download with a higher version than the running workflow.
// Install() downloads the latest version and asks Alfred to install it.
//
// Because downloading releases is slow and workflows need to run fast,
// you should not run CheckForUpdate() in a Script Filter.
//
// If an Updater is set on a Workflow struct, a magic action will be set for
// updates, so you can just add an Item that autocompletes to the update
// magic argument ("workflow:update" by default), and AwGo will check for an
// update and install it if available.
//
// See ../examples/update for a full example implementation of updates.
type Updater struct {
	Source         Source // Provides downloads
	CurrentVersion SemVer // Version of the installed workflow
	Prereleases    bool   // Include pre-releases when checking for updates

	// AlfredVersion is the version of the running Alfred application.
	// Read from $alfred_version environment variable.
	AlfredVersion SemVer

	// When the remote release list was last checked (and possibly cached)
	LastCheck      time.Time
	updateInterval time.Duration // How often to check for an update
	downloads      []Download    // Available workflow files

	// Cache paths
	cacheDir      string // Directory to store cache files in
	pathLastCheck string // Cache path for check time
	pathDownloads string // Cache path for available downloads
}

// NewUpdater creates a new Updater for Source. `currentVersion` is the workflow's
// version number and `cacheDir` is a directory where the Updater can cache
// a list of available releases.
func NewUpdater(src Source, currentVersion, cacheDir string) (*Updater, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load LastCheck

// UpdateAvailable returns true if an update is available. Retrieves
// the list of releases from the cache written by CheckForUpdate.
func (u *Updater) UpdateAvailable() bool { _ = "STUB: not implemented"; return false }

// CheckDue returns true if the time since the last check is greater than
// Updater.UpdateInterval.
func (u *Updater) CheckDue() bool { _ = "STUB: not implemented"; return false }

// log.Println("never checked for updates")

// CheckForUpdate fetches the list of releases from remote (via Releaser)
// and caches it locally.
func (u *Updater) CheckForUpdate() error {
	_ = "STUB: not implemented"
	// If update fails, don't try again for at least an hour
	return nil
}

// Install downloads and installs the latest available version.
// After the workflow file is downloaded, Install calls Alfred to
// install the update.
func (u *Updater) Install() error { _ = "STUB: not implemented"; return nil }

// clearCache removes the update cache.
func (u *Updater) clearCache() { _ = "STUB: not implemented"; return }

// cacheLastCheck saves time to cache.
func (u *Updater) cacheLastCheck() { _ = "STUB: not implemented"; return }

// Returns latest version that is compatible with the Updater's
// Alfred version & pre-release preference.
func (u *Updater) latest() *Download { _ = "STUB: not implemented"; return nil }

// Load from cache

// // Mockable function to run commands
// type commandRunner func(name string, arg ...string) error
//
// // Run command via exec.Command
// func runCommand(name string, arg ...string) error {
// 	return exec.Command(name, arg...).Run()
// }

// makeHTTPClient returns an http.Client with a sensible configuration.
func makeHTTPClient() *http.Client { _ = "STUB: not implemented"; return nil }

// getURL returns the contents of a URL.
func getURL(url string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// openURL returns an http.Response. It will return an error if the
// HTTP status code > 299.
func openURL(url string) (*http.Response, error) { _ = "STUB: not implemented"; return nil, nil }
