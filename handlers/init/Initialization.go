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
	"github.com/socifi/go-logging-facility/handlers/text"
	"github.com/tj/go-elastic"
)

type LogConfig struct {
	LogLevel string      `json:"logLevel"`
	Handlers interface{} `json:"handlers,omitempty"`
	Context  log.Fields  `json:"context"`
	Env      string      `json:"env"`
	Project  string      `json:"project"`
}

func Init(config LogConfig) (logger *log.Entry) {
	h, _ := config.Handlers.(map[string]interface{})

	var handlers []log.Handler
	if h["json"] != nil {
		info, _ := h["json"].(map[string]string)
		var file *os.File
		if info["file"] == "stdout" || info["file"] == "" {
			file = os.Stdout
		} else {
			file, _ = os.Open(info["file"])
		}
		handlers = append(handlers, json.New(file))
	}

	if h["text"] != nil {
		info, _ := h["text"].(map[string]string)
		var file *os.File
		if info["file"] == "stdout" || info["file"] == "" {
			file = os.Stdout
		} else {
			file, _ = os.Open(info["file"])
		}
		handlers = append(handlers, text.New(file))
	}

	if h["elastic"] != nil {
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
