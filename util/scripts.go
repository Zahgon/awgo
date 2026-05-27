// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package util

import (
	"errors"
	"os/exec"
)

// ErrUnknownFileType is returned by Run for files it can't identify.
var ErrUnknownFileType = errors.New("unknown filetype")

// Default Runners used by Run to determine how to execute a file.
var (
	Executable Runner // run executable files directly
	Script     Runner // run script files with commands from Interpreters

	// DefaultInterpreters maps script file extensions to interpreters.
	// Used by the Script Runner (and by extension Run()) to determine
	// how to run files that aren't executable.
	DefaultInterpreters = map[string][]string{
		".py":          {"/usr/bin/python"},
		".rb":          {"/usr/bin/ruby"},
		".sh":          {"/bin/bash"},
		".zsh":         {"/bin/zsh"},
		".scpt":        {"/usr/bin/osascript"},
		".scptd":       {"/usr/bin/osascript"},
		".applescript": {"/usr/bin/osascript"},
		".js":          {"/usr/bin/osascript", "-l", "JavaScript"},
	}

	// Available runners in order they should be tried.
	// Executable and Script are added by init.
	runners Runners
)

func init() {
	// Default runners
	Executable = &ExecRunner{}
	Script = NewScriptRunner(DefaultInterpreters)

	runners = Runners{
		Executable,
		Script,
	}
}

// Runner knows how to execute a file passed to it.
// It is used by Run to determine how to run a file.
//
// When Run is passed a filepath, it asks each registered Runner
// in turn whether it can handle the file.
type Runner interface {
	// Can Runner execute this (type of) file?
	CanRun(filename string) bool
	// Cmd that executes file (via Runner's execution mechanism).
	Cmd(filename string, args ...string) *exec.Cmd
}

// Runners implements Runner over a sequence of Runner objects.
type Runners []Runner

// CanRun returns true if one of the runners can run this file.
func (rs Runners) CanRun(filename string) bool { _ = "STUB: not implemented"; return false }

// Cmd returns a command to run the (script) file.
func (rs Runners) Cmd(filename string, args ...string) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}

// Run runs the executable or script at path and returns the output.
// If it can't figure out how to run the file (see Runner), it
// returns ErrUnknownFileType.
func (rs Runners) Run(filename string, args ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// See if a runner will accept file

// Run runs the executable or script at path and returns the output.
// If it can't figure out how to run the file (see Runner), it
// returns ErrUnknownFileType.
func Run(filename string, args ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunAS executes AppleScript and returns the output.
func RunAS(script string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RunJS executes JavaScript (JXA) and returns the output.
func RunJS(script string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// runOsaScript executes a script with /usr/bin/osascript.
// It returns the output from STDOUT.
func runOsaScript(script, lang string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Remove trailing newline added by osascript

// RunCmd executes a command and returns its output.
//
// The main difference to exec.Cmd.Output() is that RunCmd writes all
// STDERR output to the log if a command fails.
func RunCmd(cmd *exec.Cmd) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// QuoteAS converts string to an AppleScript string literal for insertion into AppleScript code.
// It wraps the value in quotation marks, so don't insert additional ones.
func QuoteAS(s string) string { _ = "STUB: not implemented"; return "" }

// QuoteJS converts a value into JavaScript source code.
// It calls json.Marshal(v), and returns an empty string if an error occurs.
func QuoteJS(v interface{}) string { _ = "STUB: not implemented"; return "" }

// ExecRunner implements Runner for executable files.
type ExecRunner struct{}

// CanRun returns true if file exists and is executable.
func (r ExecRunner) CanRun(filename string) bool { _ = "STUB: not implemented"; return false }

// Cmd returns a Cmd to run executable with args.
func (r ExecRunner) Cmd(executable string, args ...string) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}

// ScriptRunner implements Runner for the specified file extensions.
// It calls the given script with the interpreter command from Interpreters.
//
// A ScriptRunner (combined with Runners, which implements Run) is a useful
// base for adding support for running scripts to your own program.
type ScriptRunner struct {
	// Interpreters is an "extension: command" mapping of file extensions
	// to commands to invoke interpreters that can run the files.
	//
	//     Interpreters = map[string][]string{
	//         ".py": []string{"/usr/bin/python"},
	//         ".rb": []string{"/usr/bin/ruby"},
	//     }
	//
	Interpreters map[string][]string
}

// NewScriptRunner creates a new ScriptRunner for interpreters.
func NewScriptRunner(interpreters map[string][]string) *ScriptRunner {
	_ = "STUB: not implemented"
	return nil
}

// Copy over defaults

// CanRun returns true if file exists and its extension is in Interpreters.
func (r ScriptRunner) CanRun(filename string) bool { _ = "STUB: not implemented"; return false }

// Cmd returns a Cmd to run filename with its interpreter.
func (r ScriptRunner) Cmd(filename string, args ...string) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}

// any remainder of interpreter command
// path to script file
// arguments to script
