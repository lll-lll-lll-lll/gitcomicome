package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/go-git/go-git/v5/plumbing/object"
)

const (
	ExitCodeOk             = 0
	ExitCodeParseFlagError = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func NewCLI(outStream io.Writer, errStream io.Writer) *CLI {
	return &CLI{outStream: outStream, errStream: errStream}
}

func (c *CLI) Run(args []string) int {
	var (
		version bool
		filter  string
		mode    string
		url     string
		cs      object.CommitIter
	)

	flags := flag.NewFlagSet("gicom", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.Usage = func() {
		fmt.Fprintf(c.errStream, usage, "gicomi")
	}
	flags.BoolVar(&version, "version", false, "print version")
	flags.StringVar(&filter, "filter", "", "filter messages")
	flags.StringVar(&mode, "mode", "", "mode")
	flags.StringVar(&url, "url", "", "git clone url")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if version {
		fmt.Fprintf(c.errStream, "gicom version %v\n", Version)
		return ExitCodeOk
	}

	if url != "" {
		dir, git := GitCloneToTmpRepo(url, "example-clone")
		defer os.RemoveAll(dir) // clean up
		cs = AllCommitLogs(git)
	} else {
		git := SelfGitRepository()
		cs = AllCommitLogs(git)
	}

	if mode != "" {
		if filter == "" {
			fmt.Fprintf(c.errStream, "not set filter")
			return ExitCodeParseFlagError
		}
	}
	if filter != "" {
		fmt.Fprintf(c.errStream, "filter text is %s\n", filter)
		f := FilterCommits(filter, cs)
		switch mode {
		case "comment":
			g := GetComments(f)
			printJsonAny(c, g)
		case "committer":
			gc := GetCommitters(f)
			printJsonAny(c, gc)
		default:
			printJsonAny(c, f)
		}

		return ExitCodeOk
	}

	return ExitCodeOk
}

func printJsonAny[S ~[]e, e any](c *CLI, f S) {
	for _, e := range f {
		b, _ := json.Marshal(e)
		var out bytes.Buffer
		if err := json.Indent(&out, b, "", "  "); err != nil {
			fmt.Fprintf(c.errStream, "filter err %s\n", err)
		}
		fmt.Fprintf(c.errStream, "%s\n", out.String())
	}
}

const usage = `
Usage: %s [options] slug path
  
Options:
  -help or h 	 		    help
  -version            		now version
  -filter=<head comment>	head string of git comment
  -mode=<option>      		now "comment", "committer". (with filter option)
  -url=<url>				git clone url
`
