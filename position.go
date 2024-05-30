package expect

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

const unknown = "<unknown>"

type Position struct {
	Dir  string
	File string
	Line int
}

func Trace(depth int) Position {
	dir, err := os.Getwd()
	_, file, line, ok := runtime.Caller(depth)
	if err != nil {
		dir = unknown
	} else if !ok {
		file = unknown
	} else {
		temp, err := filepath.Rel(dir, file)
		if err == nil {
			file = temp
		}
	}

	return Position{Dir: dir, File: file, Line: line}
}

func (p Position) String() string { return p.File + ":" + strconv.Itoa(p.Line) }
