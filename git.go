package main

import (
	"fmt"
	"os"
	"strings"
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

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
