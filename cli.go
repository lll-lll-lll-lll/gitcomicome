package main

import (
	"flag"
	"fmt"
	"io"
)

// 終了コード
const (
	ExitCodeOk             = 0
	ExitCodeParseFlagError = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var version bool
	flags := flag.NewFlagSet("gicom", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "print version")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	if version {
		fmt.Fprintf(c.errStream, "gicom --version %v\n", Version)
		return ExitCodeOk
	}
	fmt.Fprintf(c.outStream, "work\n")
	return ExitCodeOk
}
