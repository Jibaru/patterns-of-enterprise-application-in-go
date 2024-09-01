package domain

// Plugin defines the interface that all plugins must implement
type Plugin interface {
	// Name returns the name of the plugin, used to identify it
	Name() string
	// Process formats the input data as per the plugin's implementation
	Process(data map[string]string) (string, error)
}
