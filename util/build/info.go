// Copyright (c) 2019 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT

package build

import (
	"os"
	"regexp"
)

// extract major version number from version string
var rxVersion = regexp.MustCompile(`^\d+`)

var (
	// Alfred's standard preferences folder, which is where the preferences bundle
	// is stored when the user isn't syncing their settings between machines
	defaultSyncDirV3 = os.ExpandEnv("${HOME}/Library/Application Support/Alfred 3")
	defaultSyncDirV4 = os.ExpandEnv("${HOME}/Library/Application Support/Alfred")
)

// Option configures Info created by New.
type Option func(info *Info)

// LibDir tells New to search a specific directory for Alfred config files.
// Default is ~/Library.
func LibDir(dir string) Option { _ = "STUB: not implemented"; return *new(Option) }

// InfoPlist tells New to parse a specific info.plist file. Default is ./info.plist.
func InfoPlist(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Info contains information about a workflow and Alfred.
//
// The information is extracted from environment variables,
// the workflow's info.plist, Alfred's own configuration files
// and finally, some defaults.
type Info struct {
	// Workflow info read from environment variables/info.plist
	Name     string // Workflow name
	Version  string // Workflow version
	BundleID string // Workflow bundle ID

	// Workflow directories
	CacheDir string // Workflow cache directory
	DataDir  string // Workflow data directory
	// Where workflow should be installed. This is
	// Alfred's workflow directory (AlfredWorkflowDir)
	// plus workflow's bundle ID.
	InstallDir string

	// Alfred info
	AlfredMajorVersion int    // Alfred's major version number
	AlfredSyncDir      string // Path of Alfred's syncfolder
	AlfredPrefsBundle  string // Path to the Alfred.alfredpreferences bundle
	AlfredWorkflowDir  string // Directory workflows are stored in
	AlfredCacheDir     string // Root directory for all workflow cache data
	AlfredDataDir      string // Root directory for all persistent workflow data

	// Directory searched for preferences files.
	// Default is ~/Library.
	dir string
	// Path to workflow's info.plist.
	// Default is ./info.plist
	ipPath string
}

// NewInfo creates a new Info. Workflow info is read from Alfred environment
// variables (if set), and from info.plist in the working directory and
// Alfred's configuration files. These paths may be changed using the
// the LibDir and InfoPlist Options. Settings from info.plist take priority
// over those from environment variables.
//
// It returns an error if info.plist or the configuration files cannot be found.
func NewInfo(option ...Option) (*Info, error) { _ = "STUB: not implemented"; return nil, nil }

// Env returns an Alfred-like environment.
func (info *Info) Env() map[string]string { _ = "STUB: not implemented"; return nil }

func (info *Info) findFolders() error { _ = "STUB: not implemented"; return nil }

func (info *Info) findAlfredVersion() error { _ = "STUB: not implemented"; return nil }

func (info *Info) readEnv() { _ = "STUB: not implemented"; return }

// readPlist reads workflow information from info.plist.
func (info *Info) readPlist() error { _ = "STUB: not implemented"; return nil }

// expand ~ in a filepath.
func expand(path string) string { _ = "STUB: not implemented"; return "" }

// get path to Alfred's sync folder (parent of Alfred.alfredpreferences) from
// environment or Alfred's config files
func findSyncFolder(v int, dir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Alfred 4+ has a dedicated prefs.json file, but earlier versions store
// the setting in Alfred Preference's version-specific prefs file

// Look for Alfred 4+ prefs.json

// Look for Alfred 3 preferences plist
