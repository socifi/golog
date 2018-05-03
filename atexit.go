package log

import (
	"fmt"
	"os"
)

const (
	// Version is package version
	Version = "0.1.0"
)

var exitHandlers = []func(){}

func runHandler(handler func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, "error: atexit handler error:", err)
		}
	}()

	handler()
}

func runHandlers() {
	for _, handler := range exitHandlers {
		runHandler(handler)
	}
}

// Exit runs all the exitHandlers and then terminates the program using
// os.Exit(code)
func Exit(code int) {
	runHandlers()
	os.Exit(code)
}

func (e *Entry) Exit(code int) {
	Exit(code)
}

func (l *Logger) Exit(code int) {
	Exit(code)
}

// Register adds a handler, call Exit in this module to invoke all exitHandlers.
func AddExitHandler(handler func()) {
	exitHandlers = append(exitHandlers, handler)
}

func (e *Entry) AddExitHandler(handler func()) {
	AddExitHandler(handler)
}

func (l *Logger) AddExitHandler(handler func()) {
	AddExitHandler(handler)
}
