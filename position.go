package expect

import (
	"os"
	"runtime"
	"strconv"
)

const unknown = "<unknown>"

type Position struct {
	WorkingDirectory string
	File             string
	Line             int
}

func Trace(depth int) Position {
	workingDir, err := os.Getwd()
	if err != nil {
		workingDir = unknown
	}
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = unknown
	} else if workingDir != unknown {
		// fixme might contains bug.
		file = file[len(workingDir)+1:]
	}
	return Position{File: file, Line: line}
}

func (p Position) String() string { return p.File + ":" + strconv.Itoa(p.Line) }
