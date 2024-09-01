package main

import (
	"fmt"

	"github.com/jibaru/plugin/internal/plugins"
	"github.com/jibaru/plugin/internal/processor"
)

// Main entry point of the application
func main() {
	// Create a processor
	processor := processor.NewProcessor()

	// Register plugins
	processor.RegisterPlugin(plugins.NewJSONPlugin())
	processor.RegisterPlugin(plugins.NewXMLPlugin())

	// Use the processor to process data in different formats
	data := map[string]string{"key": "value"}
	formattedData, err := processor.ProcessData("json", data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Processed data (JSON):", formattedData)

	formattedData, err = processor.ProcessData("xml", data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Processed data (XML):", formattedData)
}
