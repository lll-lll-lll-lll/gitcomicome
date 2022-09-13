package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func TestGitClone(t *testing.T) {
	url := "https://github.com/lll-lll-lll-lll/webvtt-reader"
	Info("git clone" + url)

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})

	CheckIfError(err)
	// Gets the HEAD history from HEAD, just like this command:
	Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	today := time.Now()
	y := today.AddDate(0, 0, -3)

	commits := Commits(r, y, today, ref)

	// ... just iterates over the commits, printing it
	err = commits.ForEach(func(c *object.Commit) error {
		fmt.Println("commit")
		fmt.Println(c)

		return nil
	})
	CheckIfError(err)
}

func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func Commits(r *git.Repository, start, until time.Time, ref *plumbing.Reference) object.CommitIter {
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &start, Until: &until})
	CheckIfError(err)
	return cIter
}
