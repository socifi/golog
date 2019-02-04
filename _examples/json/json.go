package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/socifi/golog"
	"github.com/socifi/golog/handler/json"
)

func one(ctx log.Interface) {
	ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	s := debug.Stack()
	fmt.Println(string(s))
}

func main() {
	log.SetHandler(json.New(os.Stderr))

	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	for range time.Tick(time.Millisecond * 200) {
		ctx.Info("upload")
		ctx.Info("upload complete")
		ctx.Warn("upload retry")
		one(ctx)
		// ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	}
}
