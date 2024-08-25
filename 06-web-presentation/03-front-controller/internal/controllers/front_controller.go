package controllers

import (
	"net/http"
)

// Command interface that all commands should implement
type Command interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

// FrontController struct that holds the registered commands
type FrontController struct {
	commands map[string]Command
}

// NewFrontController creates a new FrontController
func NewFrontController() *FrontController {
	return &FrontController{
		commands: make(map[string]Command),
	}
}

// RegisterCommand registers a command with a specific path
func (fc *FrontController) RegisterCommand(path string, command Command) {
	fc.commands[path] = command
}

// ServeHTTP handles all incoming HTTP requests
func (fc *FrontController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	command, exists := fc.commands[path]
	if !exists {
		http.NotFound(w, r)
		return
	}
	command.Execute(w, r)
}
