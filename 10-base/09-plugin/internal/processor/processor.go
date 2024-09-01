package processor

import (
	"errors"

	"github.com/jibaru/plugin/internal/domain"
)

// Processor handles processing data using registered plugins
type Processor struct {
	plugins map[string]domain.Plugin
}

// NewProcessor creates a new Processor with no plugins
func NewProcessor() *Processor {
	return &Processor{
		plugins: make(map[string]domain.Plugin),
	}
}

// RegisterPlugin adds a new plugin to the processor
func (p *Processor) RegisterPlugin(plugin domain.Plugin) {
	p.plugins[plugin.Name()] = plugin
}

// ProcessData processes data using a specified plugin by name
func (p *Processor) ProcessData(pluginName string, data map[string]string) (string, error) {
	plugin, exists := p.plugins[pluginName]
	if !exists {
		return "", errors.New("plugin not found")
	}
	return plugin.Process(data)
}
