package log

import (
	"fmt"
	"os"
	"sync"
)

var exitHandlers = []func(){}
var mux = &sync.Mutex{}

// runHandler tries to run exactly one handler and if not successful, writes error message to stderr
func runHandler(handler func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, "error: atexit handler error:", err)
		}
	}()

	handler()
}

// runHandlers is a helper function that runs all functions registered in exitHandlers
func runHandlers() {
	for _, handler := range exitHandlers {
		runHandler(handler)
	}
}

// Exit runs all the exitHandlers and then terminates the program using
// os.Exit(code)
func Exit(code int) {
	mux.Lock()
	defer mux.Unlock()
	runHandlers()
	os.Exit(code)
}

// Exit runs all the exitHandlers and then terminates the program using
// os.Exit(code)
func (e *Entry) Exit(code int) {
	Exit(code)
}

// Exit runs all the exitHandlers and then terminates the program using
// os.Exit(code)
func (l *Logger) Exit(code int) {
	Exit(code)
}

// AddExitHandler adds a handler, call Exit in this module to invoke all exitHandlers.
// Thread safe
func AddExitHandler(handler func()) {
	mux.Lock()
	defer mux.Unlock()
	exitHandlers = append(exitHandlers, handler)
}

// AddExitHandler adds a handler, call Exit in this module to invoke all exitHandlers.
// Thread safe
func (e *Entry) AddExitHandler(handler func()) {
	AddExitHandler(handler)
}

// AddExitHandler adds a handler, call Exit in this module to invoke all exitHandlers.
// Thread safe
func (l *Logger) AddExitHandler(handler func()) {
	AddExitHandler(handler)
}
