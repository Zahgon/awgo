// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package aw

import (
	"go.deanishe.net/env"
)

// To populates (tagged) struct v with values from the environment.
func (cfg *Config) To(v interface{}) error { _ = "STUB: not implemented"; return nil }

// From saves the fields of (tagged) struct v to the workflow's settings in Alfred.
// All supported and unignored fields are saved by default. The behaviour can be
// customised by passing in options from deanishe/go-env, such as env.IgnoreZeroValues
// to omit any fields set to zero values.
//
// https://godoc.org/go.deanishe.net/env#DumpOption
func (cfg *Config) From(v interface{}, opt ...env.DumpOption) error {
	_ = "STUB: not implemented"
	return nil
}

// setMulti batches the saving of multiple variables.
func (cfg *Config) setMulti(variables map[string]string, export bool) error {
	_ = "STUB: not implemented"
	// sort keys to make the output testable
	return nil
}
