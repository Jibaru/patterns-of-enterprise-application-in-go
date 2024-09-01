package plugins

import (
	"encoding/xml"

	"github.com/jibaru/plugin/internal/domain"
)

// XMLPlugin is a plugin that processes data into XML format
type XMLPlugin struct{}

// NewXMLPlugin creates a new XMLPlugin instance
func NewXMLPlugin() domain.Plugin {
	return &XMLPlugin{}
}

// Name returns the name of the plugin
func (p *XMLPlugin) Name() string {
	return "xml"
}

// KeyValue is a helper struct to represent map data in a key-value format for XML marshaling
type KeyValue struct {
	XMLName xml.Name `xml:"item"`
	Key     string   `xml:"key"`
	Value   string   `xml:"value"`
}

// Process converts the input data to XML format
func (p *XMLPlugin) Process(data map[string]string) (string, error) {
	// Convert the map to a slice of KeyValue pairs
	keyValuePairs := make([]KeyValue, 0, len(data))
	for key, value := range data {
		keyValuePairs = append(keyValuePairs, KeyValue{Key: key, Value: value})
	}

	// Marshal the slice of KeyValue pairs to XML
	xmlData, err := xml.MarshalIndent(keyValuePairs, "", "  ")
	if err != nil {
		return "", err
	}

	return string(xmlData), nil
}
