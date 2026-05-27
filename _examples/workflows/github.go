// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package main

import (
	"net"
	"net/http"
	"time"
)

var (
	apiURL     = "https://api.github.com/search/repositories?per_page=100"
	apiQuery   = "topic:alfred-workflow"
	apiHeaders = map[string]string{
		"Accept":     "application/vnd.github.mercy-preview+json",
		"User-Agent": "github.com/deanishe/awgo",
	}
	client *http.Client
)

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 60 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   30 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
			ExpectContinueTimeout: 10 * time.Second,
		},
	}
}

// Repo is a GitHub repo from the GitHub API.
type Repo struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Owner       *repoOwner `json:"owner"`
	URL         string     `json:"html_url"`
	Stars       int64      `json:"stargazers_count"`
	Topics      []string   `json:"topics"`
	Lang        string     `json:"language"`
}

// FullName returns standard "owner/repo" format.
func (r *Repo) FullName() string { _ = "STUB: not implemented"; return "" }

// Username is GitHub user login.
func (r *Repo) Username() string { _ = "STUB: not implemented"; return "" }

// repoOwner is a helper struct for unmarshalling API JSON response.
type repoOwner struct {
	Login string `json:"login"`
}

// apiResponse is the top-level helper struct for unmarshalling API JSON response.
type apiResponse struct {
	Repos []*Repo `json:"items"`
	Total int     `json:"total_count"`
}

// fetchRepos fetches all repos with topic "alfred-workflow" from GitHub.
//
// It iterates through all pages of results, returning all matching repos.
func fetchRepos() ([]*Repo, error) { _ = "STUB: not implemented"; return nil, nil }

// Generate URL for next page of results

// Add headers to request (label feature isn't part of the standard API yet)

// Parse response

// Populate pageCount if unset
