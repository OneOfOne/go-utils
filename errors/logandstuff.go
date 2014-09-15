package errors

import (
	"fmt"
	"os"
)

// DieIf prints msg + err to stderr then exits the program with status code 1
// it's like a panic, without the stacktrace.
// use %v in message to print the error, default is "error: %v"
func DieIf(err error, msg string) {
	if err != nil {
		if len(msg) == 0 {
			msg = "error: %v"
		}
		fmt.Fprintf(os.Stderr, msg+"\n", err)
		os.Exit(1)
	}
}

// PrintfIf calls fn (log.Printf or fmt.Printf for example) with msg and err
// if err != nil, it will return true and call fn(msg, err)
// example:
//	if PrintfIf(err, "error: %v\n", log.Printf) {
//		return
//	}
func PrintfIf(err error, msg string, fn func(format string, args ...interface{})) (b bool) {
	if b = err != nil; b {
		if len(msg) == 0 {
			msg = "error: %v"
		}
		fn(msg, err)
	}
	return
}
