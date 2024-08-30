package controllers

import (
	"net/http"
)

type InputController struct {
	AppController *ApplicationController
}

func (ic *InputController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Path[len("/action/"):]
	domainCommand := ic.AppController.GetDomainCommand(action, r)
	if domainCommand != nil {
		result := domainCommand.Run()
		viewFunc := ic.AppController.GetView(action)
		if viewFunc != nil {
			viewFunc(w, result)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}
