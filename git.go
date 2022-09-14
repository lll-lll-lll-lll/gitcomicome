package main

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func SpanCommitLogs(r *git.Repository, start, until time.Time, ref *plumbing.Reference) object.CommitIter {
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &start, Until: &until})
	CheckIfError(err)
	return cIter
}

func SelfGitRepository() *git.Repository {
	s, err := git.PlainOpen("./")
	if err != nil {
		CheckIfError(err)
	}
	return s
}

func SelfAllCommitLogs(s *git.Repository) object.CommitIter {
	c, err := s.CommitObjects()
	if err != nil {
		CheckIfError(err)
	}
	return c
}
