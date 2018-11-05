package json_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/socifi/go-logging-facility"
	"github.com/socifi/go-logging-facility/handler/json"
)

func init() {
	log.Now = func() time.Time {
		return time.Unix(0, 0).UTC()
	}
}

func Test(t *testing.T) {
	var buf bytes.Buffer

	log.SetHandler(json.New(&buf))
	log.WithField("user", "tj").WithField("id", "123").Info("hello")
	log.Info("world")
	log.Error("boom")

	expected := `{"context":{"id":"123","user":"tj"},"level":200,"level_name":"info","timestamp":"1970-01-01T00:00:00Z","message":"hello","env":"","project":"","hostname":""}
{"context":{},"level":200,"level_name":"info","timestamp":"1970-01-01T00:00:00Z","message":"world","env":"","project":"","hostname":""}
{"context":{},"level":400,"level_name":"error","timestamp":"1970-01-01T00:00:00Z","message":"boom","env":"","project":"","hostname":""}
`

	assert.Equal(t, expected, buf.String())
}
