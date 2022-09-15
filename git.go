package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
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

func loadJsonWithValidate(i string) (interface{}, error) {
	b, _ := ioutil.ReadFile(i)
	var j interface{}
	_ = json.Unmarshal(b, &j)
	r := j.(map[string]interface{})["Rule"]
	if r == nil {
		return nil, errors.New("error not have `Rule` property")
	}
	d := r.(map[string]interface{})["Separator"]
	if d == nil {
		return nil, errors.New("error not have `Separator` property in `Rule` property")
	}
	return r, nil
}

func FilterCommits(filter string, c object.CommitIter) []*object.Commit {
	var s []*object.Commit
	c.ForEach(func(c *object.Commit) error {
		b := CheckRegexp(fmt.Sprintf("^%s", filter), c.Message)
		if b {
			s = append(s, c)
		}
		return nil
	})
	return s
}

func GetComments(c []*object.Commit) []string {
	var s []string
	for _, x := range c {
		s = append(s, x.Message)
	}
	return s
}

func GetCommitters(c []*object.Commit) []object.Signature {
	var s []object.Signature
	for _, x := range c {
		s = append(s, x.Committer)
	}
	return s
}

func CheckRegexp(pattern, str string) bool {
	m, err := regexp.Match(pattern, []byte(str))
	if err != nil {
		return false
	}
	return m
}
