// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package aw

import (
	"os/exec"
)

// ErrJobExists is the error returned by RunInBackground if a job with
// the given name is already running.
type ErrJobExists struct {
	Name string // Name of the job
	Pid  int    // PID of the running job
}

// Error implements error interface.
func (err ErrJobExists) Error() string { _ = "STUB: not implemented"; return "" }

// Is returns true if target is of type ErrJobExists.
func (err ErrJobExists) Is(target error) bool { _ = "STUB: not implemented"; return false }

// IsJobExists returns true if error is of type or wraps ErrJobExists.
func IsJobExists(err error) bool { _ = "STUB: not implemented"; return false }

// RunInBackground executes cmd in the background. It returns an
// ErrJobExists error if a job of the same name is already running.
func (wf *Workflow) RunInBackground(jobName string, cmd *exec.Cmd) error {
	_ = "STUB: not implemented"
	return nil
}

// Prevent process from being killed when parent is

// Kill stops a background job.
func (wf *Workflow) Kill(jobName string) error { _ = "STUB: not implemented"; return nil }

// IsRunning returns true if a job with name jobName is currently running.
func (wf *Workflow) IsRunning(jobName string) bool { _ = "STUB: not implemented"; return false }

// Delete stale PID file

// Save PID to a job-specific file.
func (wf *Workflow) savePid(jobName string, pid int) error { _ = "STUB: not implemented"; return nil }

// Return PID for job.
func (wf *Workflow) getPid(jobName string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Path to PID file for job.
func (wf *Workflow) pidFile(jobName string) string { _ = "STUB: not implemented"; return "" }
