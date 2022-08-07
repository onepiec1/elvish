//go:build !windows && !plan9 && !js

package main

import (
	"strings"
)

type opType int

// Operations for driving a demo ttyshot.
const (
	opText   opType = iota // send the provided text, optionally followed by Enter
	opPrompt               // wait for prompt marker
	opTmux                 // run tmux command
)

type op struct {
	typ opType
	val any
}

func parseSpec(content string) []op {
	lines := strings.Split(content, "\n")
	ops := make([]op, 1, len(lines)+2)
	ops[0] = op{opPrompt, nil}

	for _, line := range lines {
		if len(line) == 0 {
			continue // ignore empty lines
		}
		var newOp op
		if line == "#prompt" {
			newOp = op{opPrompt, nil}
		} else if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "# ") {
			newOp = op{opTmux, strings.Fields(line[1:])}
		} else {
			newOp = op{opText, line}
		}
		ops = append(ops, newOp)
	}

	if len(ops) > 0 && ops[len(ops)-1].typ == opText {
		// The termination of the ttyshot process relies on the editor to be
		// active, so add an implicit #prompt.
		ops = append(ops, op{opPrompt, nil})
	}

	return ops
}