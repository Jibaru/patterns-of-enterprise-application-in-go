package plugins

import (
	"encoding/json"

	"github.com/jibaru/plugin/internal/domain"
)

// JSONPlugin is a plugin that processes data into JSON format
type JSONPlugin struct{}

// NewJSONPlugin creates a new JSONPlugin instance
func NewJSONPlugin() domain.Plugin {
	return &JSONPlugin{}
}

// Name returns the name of the plugin
func (p *JSONPlugin) Name() string {
	return "json"
}

// Process converts the input data to JSON format
func (p *JSONPlugin) Process(data map[string]string) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
