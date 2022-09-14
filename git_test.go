package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestGitClone(t *testing.T) {
	url := "https://github.com/lll-lll-lll-lll/webvtt-reader"
	Info("git clone" + url)
	dir, err := ioutil.TempDir(".", "clone-example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up
	fmt.Print(dir)

	r, err := git.PlainClone(dir, false, &git.CloneOptions{
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

	commits := SpanCommitLogs(r, y, today, ref)

	// ... just iterates over the commits, printing it
	err = commits.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	CheckIfError(err)
}

func TestPlainSelfOpen(t *testing.T) {
	t.Run("open self .git dir", func(t *testing.T) {
		s, err := git.PlainOpen("./")
		if err != nil {
			t.Log(err)
		}
		a, err := s.CommitObjects()
		if err != nil {
			t.Log(err)
		}
		a.ForEach(func(c *object.Commit) error {
			t.Log(c.Message)
			return nil
		})
	})
	t.Run("load json function", func(t *testing.T) {
		j, err := LoadJsoWithnValidate("./rule.json")
		if err != nil {
			t.Error(err)
		}
		t.Log(j)
	})
}
