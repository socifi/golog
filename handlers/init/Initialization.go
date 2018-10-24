package loginit

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/socifi/go-logging-facility"
	"github.com/socifi/go-logging-facility/handlers/es"
	"github.com/socifi/go-logging-facility/handlers/json"
	"github.com/socifi/go-logging-facility/handlers/multi"
	"github.com/tj/go-elastic"
)

// LogConfig contains all needed information for logger initialization
type LogConfig struct {
	LogLevel string      `json:"logLevel"`
	Handlers interface{} `json:"handlers,omitempty"`
	Context  log.Fields  `json:"context"`
	Env      string      `json:"env"`
	Project  string      `json:"project"`
}

// Init initializes logger with values from LogConfig structure
func Init(config LogConfig) (logger *log.Entry) {
	h, _ := config.Handlers.(map[string]interface{})

	var handlers []log.Handler
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

	log.SetHandler(multi.New(handlers...))
	log.SetLevelFromString(config.LogLevel)
	log.WithFields(config.Context)

	logger = log.SetEnvProject(config.Env, config.Project)
	logger.Debug("Logger initialized")
	return logger
}
