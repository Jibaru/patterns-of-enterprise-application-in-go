package storage

// Storage defines an interface for saving and retrieving data
type Storage interface {
	Save(key string, value string) error
	Load(key string) (string, error)
}
