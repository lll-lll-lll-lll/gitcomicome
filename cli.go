package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	)
	git := SelfGitRepository()
	cs := SelfAllCommitLogs(git)
	flags := flag.NewFlagSet("gicom", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "print version")
	flags.StringVar(&filter, "filter", "", "filter messages")
	flags.StringVar(&mode, "mode", "", "mode")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	if version {
		fmt.Fprintf(c.errStream, "gicom version %v\n", Version)
		return ExitCodeOk
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
