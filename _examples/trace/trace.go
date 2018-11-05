package main

import (
	"os"
	"time"

	"github.com/socifi/go-logging-facility"
	"github.com/socifi/go-logging-facility/handler/text"
)

func work(ctx log.Interface) (err error) {
	path := "Readme.md"
	defer ctx.WithField("path", path).Trace("opening").Stop(&err)
	_, err = os.Open(path)
	return
}

func main() {
	log.SetHandler(text.New(os.Stderr))

	ctx := log.WithFields(log.Fields{
		"app": "myapp",
		"env": "prod",
	})

	for range time.Tick(time.Second) {
		_ = work(ctx)
	}
}
