package main

import (
	"fmt"
	"log"

	"github.com/jibaru/separated-interface/internal/memory"
	"github.com/jibaru/separated-interface/internal/sqlite"
	"github.com/jibaru/separated-interface/internal/storage"
)

func useStorage(store storage.Storage) {
	// Save some data
	err := store.Save("username", "johndoe")
	if err != nil {
		log.Fatalf("failed to save data: %v", err)
	}
	fmt.Println("Saved value:", "johndoe")

	// Load the data
	value, err := store.Load("username")
	if err != nil {
		log.Fatalf("failed to load data: %v", err)
	}

	fmt.Println("Loaded value:", value)
}

func main() {
	var store storage.Storage
	var err error

	store, err = sqlite.NewSQLiteStorage("data.db")
	if err != nil {
		log.Fatalf("failed to initialize SQLite storage: %v", err)
	}

	fmt.Println("Using sqlite storage")
	useStorage(store)

	store = memory.NewMemoryStorage()

	fmt.Println("Using memory storage")
	useStorage(store)
}
