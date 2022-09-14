package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

func LoadJsoWithnValidate(i string) (interface{}, error) {
	b, _ := ioutil.ReadFile(i)
	var j interface{}
	_ = json.Unmarshal(b, &j)
	a := j.(map[string]interface{})["Rule"]
	if a == nil {
		return nil, errors.New("error not have `Rule` property")
	}
	return a, nil
}
