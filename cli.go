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
	)
	git := SelfGitRepository()
	cs := SelfAllCommitLogs(git)
	flags := flag.NewFlagSet("gicom", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "print version")
	flags.StringVar(&filter, "filter", "", "filter messages")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	if version {
		fmt.Fprintf(c.errStream, "gicom version %v\n", Version)
		return ExitCodeOk
	}
	if filter != "" {
		fmt.Fprintf(c.errStream, "filter text is %s\n", filter)
		f := FilterCommits(filter, cs)
		for _, e := range f {
			b, _ := json.Marshal(e)
			var out bytes.Buffer
			if err := json.Indent(&out, b, "", "  "); err != nil {
				fmt.Fprintf(c.errStream, "filter err %s\n", err)
				return ExitCodeParseFlagError
			}
			fmt.Fprintf(c.errStream, "%s\n", out.String())
		}

		return ExitCodeOk
	}

	return ExitCodeOk
}
