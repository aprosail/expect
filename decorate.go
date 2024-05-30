package expect

import (
	"bytes"
)

func dim(message string) string    { return "\x1b[2m" + message + "\x1b[0m" }
func red(message string) string    { return "\x1b[31m" + message + "\x1b[0m" }
func green(message string) string  { return "\x1b[32m" + message + "\x1b[0m" }
func yellow(message string) string { return "\x1b[33m" + message + "\x1b[0m" }

func backspace(times int) string {
	var buffer bytes.Buffer
	for i := 0; i < times; i++ {
		buffer.WriteByte('\b')
	}
	return buffer.String()
}
