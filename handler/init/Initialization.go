package log

import (
	"fmt"
	"net/http"
	"os"
	"time"

	lg "github.com/socifi/go-logging-facility"
//	"github.com/socifi/go-logging-facility/handler/discard"
	"github.com/socifi/go-logging-facility/handler/es"
	"github.com/socifi/go-logging-facility/handler/json"
	"github.com/socifi/go-logging-facility/handler/multi"
	"github.com/tj/go-elastic"
)

type Log lg.Interface

// LogConfig contains all needed information for logger initialization
type Config struct {
	LogLevel string      `json:"logLevel"`
	Handlers interface{} `json:"handlers,omitempty"`
	Context  lg.Fields  `json:"context"`
	Env      string      `json:"env"`
	Project  string      `json:"project"`
}

// Init initializes logger with values from LogConfig structure
func Init(config Config) (logger Log) {
	h, _ := config.Handlers.(map[string]interface{})

	var handlers []lg.Handler
	if (h["json"]) != nil {
		info, _ := h["json"].(map[string]string)
		var file *os.File
		if info["file"] == "stdout" || info["file"] == "" {
			file = os.Stdout
		} else {
			file, _ = os.Open(info["file"])
		}
		handlers = append(handlers, json.New(file))
	}

	if (h["elastic"]) != nil {
		info, _ := h["elastic"].(map[string]interface{})
		url := fmt.Sprintf("%s:%d", info["host"].(string), int(info["port"].(float64)))
		esClient := elastic.New(url)
		esClient.HTTPClient = &http.Client{Timeout: time.Duration(info["timeout"].(float64)) * time.Second}

		handlers = append(handlers, es.New(&es.Config{
			Client:     esClient,
			Format:     info["format"].(string),
			Type:       info["type"].(string),
			BufferSize: 1,
		}))
	}

	lg.SetHandler(multi.New(handlers...))
	lg.SetLevelFromString(config.LogLevel)
	lg.WithFields(config.Context)

	logger = lg.SetEnvProject(config.Env, config.Project)
	logger.Debug("Logger initialized")
	return logger
}
