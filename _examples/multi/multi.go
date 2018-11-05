package main

import (
	"errors"
	"os"
	"time"

	"github.com/socifi/go-logging-facility"
	"github.com/socifi/go-logging-facility/handler/json"
	"github.com/socifi/go-logging-facility/handler/multi"
	"github.com/socifi/go-logging-facility/handler/text"
)

func main() {
	log.SetHandler(multi.New(
		text.New(os.Stderr),
		json.New(os.Stderr),
	))

	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	for range time.Tick(time.Millisecond * 200) {
		ctx.Info("upload")
		ctx.Info("upload complete")
		ctx.Warn("upload retry")
		ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	}
}
