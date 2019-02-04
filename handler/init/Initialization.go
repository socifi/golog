package log

import (
	"fmt"
	"net/http"
	"os"
	"time"

	lg "github.com/socifi/golog"
	"github.com/socifi/golog/handler/discard"
	"github.com/socifi/golog/handler/es"
	"github.com/socifi/golog/handler/json"
	"github.com/socifi/golog/handler/kinesis"
	"github.com/socifi/golog/handler/logfmt"
	"github.com/socifi/golog/handler/memory"
	"github.com/socifi/golog/handler/multi"
	"github.com/socifi/golog/handler/papertrail"
	"github.com/tj/go-elastic"
)

type Log lg.Interface

// LogConfig contains all needed information for logger initialization
type Config struct {
	LogLevel string      `json:"logLevel"`
	Handlers interface{} `json:"handlers,omitempty"`
	Context  lg.Fields   `json:"context"`
	Env      string      `json:"env"`
	Project  string      `json:"project"`
}

// initJSON initializes new JSON log with given settings
func initJSON(settings interface{}) (lg.Handler, error) {
	s, ok := settings.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Error converting json handler data to map")
	}
	var file *os.File
	var err error
	if s["file"] == "stdout" || s["file"] == "" {
		file = os.Stdout
	} else {
		file, err = os.Create(s["file"].(string))
		if err != nil {
			return nil, err
		}
	}
	return json.New(file), nil
}

// initElastic initializes new elastic log with given settings
func initElastic(settings interface{}) (lg.Handler, error) {
	s, ok := settings.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Error converting elastic handler data to map")
	}

	url := fmt.Sprintf("%s:%d", s["host"].(string), int(s["port"].(float64)))
	esClient := elastic.New(url)
	esClient.HTTPClient = &http.Client{Timeout: time.Duration(s["timeout"].(float64)) * time.Second}

	return es.New(&es.Config{
		Client:     esClient,
		Format:     s["format"].(string),
		Type:       s["type"].(string),
		BufferSize: 1,
	}), nil
}

// initDiscard initializes new discard log
func initDiscard() lg.Handler {
	return discard.New()
}

// initKinesis initializes new kinesis log with given settings
func initKinesis(settings interface{}) (lg.Handler, error) {
	s, ok := settings.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Error converting kinesis handler data to map")
	}
	return kinesis.New(s["stream"].(string)), nil
}

// initLogfmt initializes new logfmt log with given settings
func initLogfmt(settings interface{}) (lg.Handler, error) {
	s, ok := settings.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Error converting logfmt handler data to map")
	}
	var file *os.File
	var err error
	if s["file"] == "stdout" || s["file"] == "" {
		file = os.Stdout
	} else {
		file, err = os.Create(s["file"].(string))
		if err != nil {
			return nil, err
		}
	}
	return logfmt.New(file), nil
}

// initMemory initializes new in memory log
func initMemory() lg.Handler {
	return memory.New()
}

// initPapertrail initializes new Papertrail log with given settings
func initPapertrail(settings interface{}) (lg.Handler, error) {
	s, ok := settings.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Error converting papertrail handler data to map")
	}

	var c papertrail.Config

	if s["host"] == nil {
		return nil, fmt.Errorf("Papertrail handler must have specified host")
	}
	c.Host = s["host"].(string)

	if s["port"] == nil {
		return nil, fmt.Errorf("Papertrail handler must have specified port")
	}
	c.Port = int(s["port"].(float64))

	var hostname string
	var err error
	if s["hostname"] == nil {
		hostname, err = os.Hostname()
		if err != nil {
			hostname = ""
		}
		c.Hostname = hostname
	} else {
		c.Hostname = s["hostname"].(string)
	}

	if s["tag"] != nil {
		c.Tag = s["tag"].(string)
	}

	fmt.Println(c)
	return papertrail.New(&c), nil
}

// Init initializes logger with values from LogConfig structure
func Init(config Config) (Log, error) {
	h, ok := config.Handlers.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Error converting handlers to map")
	}

	var handlers []lg.Handler
	if (h["json"]) != nil {
		handler, err := initJSON(h["json"])
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
	}

	if (h["elastic"]) != nil {
		handler, err := initElastic(h["elastic"])
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
	}

	if (h["discard"]) != nil {
		handlers = append(handlers, initDiscard())
	}

	if (h["kinesis"]) != nil {
		handler, err := initKinesis(h["kinesis"])
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
	}

	if (h["logfmt"]) != nil {
		handler, err := initLogfmt(h["logfmt"])
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
	}

	if (h["memory"]) != nil {
		handlers = append(handlers, initMemory())
	}

	if (h["papertrail"]) != nil {
		handler, err := initPapertrail(h["papertrail"])
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
	}

	lg.SetHandler(multi.New(handlers...))
	lg.SetLevelFromString(config.LogLevel)
	lg.WithFields(config.Context)

	logger := lg.SetEnvProject(config.Env, config.Project)
	logger.Debug("Logger initialized")
	return logger, nil
}
