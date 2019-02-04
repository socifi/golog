package logfmt_test

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/socifi/golog"
	"github.com/socifi/golog/handler/logfmt"
)

func init() {
	log.Now = func() time.Time {
		return time.Unix(0, 0).UTC()
	}
}

func Test(t *testing.T) {
	var buf bytes.Buffer

	log.SetHandler(logfmt.New(&buf))
	log.WithField("user", "tj").WithField("id", "123").Info("hello")
	log.Info("world")
	log.Error("boom")

	expected := `timestamp=1970-01-01T00:00:00Z level=200 message=hello id=123 user=tj
timestamp=1970-01-01T00:00:00Z level=200 message=world
timestamp=1970-01-01T00:00:00Z level=400 message=boom
`

	assert.Equal(t, expected, buf.String())
}

func Benchmark(b *testing.B) {
	log.SetHandler(logfmt.New(ioutil.Discard))
	ctx := log.WithField("user", "tj").WithField("id", "123")

	for i := 0; i < b.N; i++ {
		ctx.Info("hello")
	}
}
