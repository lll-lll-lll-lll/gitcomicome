package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunVersionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("gicom --version", " ")
	status := cli.Run(args)
	if status != ExitCodeOk {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOk)
	}
	expected := fmt.Sprintf("gicom --version %s\n", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("Output=%q, want %q", errStream.String(), expected)
	}
}
