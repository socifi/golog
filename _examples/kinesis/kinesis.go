package main

import (
	"os"
	"time"

	"github.com/socifi/golog"
	"github.com/socifi/golog/handler/kinesis"
	"github.com/socifi/golog/handler/multi"
	"github.com/socifi/golog/handler/text"
)

func main() {
	log.SetHandler(multi.New(
		text.New(os.Stderr),
		kinesis.New("logs"),
	))

	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	for range time.Tick(time.Millisecond * 100) {
		ctx.Info("upload")
		ctx.Info("upload complete")
	}
}
