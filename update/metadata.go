// Copyright (c) 2019 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT

package update

import (
	aw "github.com/deanishe/awgo"
)

// Metadata is a Workflow Option. It sets a Workflow Updater based on
// a `metadata.json` file exported from Alfred 4+.
//
// URL is the location of the `metadata.json` file. Note: You *must*
// set `downloadurl` in the `metadata.json` file to the URL
// of your .alfredworkflow (or .alfred4workflow etc.) file.
func Metadata(url string) aw.Option { _ = "STUB: not implemented"; return *new(aw.Option) }

type metadataSource struct {
	url   string
	dl    *Download
	fetch func(URL string) ([]byte, error)
}

// Downloads implements Source.
func (src *metadataSource) Downloads() ([]Download, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// data model for metadata.json JSON.
type metadataRelease struct {
	Data struct {
		URL     string `json:"downloadurl"`
		Version string `json:"version"`
	} `json:"alfredworkflow"`
}

func parseMetadata(data []byte) (Download, error) {
	_ = "STUB: not implemented"
	return *new(Download), nil
}
