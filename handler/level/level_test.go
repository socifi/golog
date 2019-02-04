package level_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/socifi/golog"
	"github.com/socifi/golog/handler/level"
	"github.com/socifi/golog/handler/memory"
)

func Test(t *testing.T) {
	h := memory.New()

	ctx := log.Logger{
		Handler: level.New(h, log.ErrorLevel),
		Level:   log.InfoLevel,
	}

	ctx.Info("hello")
	ctx.Info("world")
	ctx.Error("boom")

	assert.Len(t, h.Entries, 1)
	assert.Equal(t, h.Entries[0].Message, "boom")
}
