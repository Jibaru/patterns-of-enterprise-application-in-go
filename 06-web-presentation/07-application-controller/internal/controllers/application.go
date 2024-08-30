package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jibaru/application-controller/internal/commands"
	"github.com/jibaru/application-controller/internal/views"
)

var AvailableActions = []string{
	"GET /action/view-user",
}

type ApplicationController struct{}

func (ac *ApplicationController) GetDomainCommand(action string, r *http.Request) commands.DomainCommand {
	switch action {
	case "view-user":
		idParam := r.URL.Query().Get("id")

		userID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			userID = 1
			log.Printf("error %v", err)
		}
		log.Printf("GET /action/view-user?id=%v\n", userID)
		return &commands.GetUserCommand{UserID: int(userID)}
	default:
		return nil
	}
}

func (ac *ApplicationController) GetView(action string) func(http.ResponseWriter, interface{}) {
	switch action {
	case "view-user":
		return views.RenderUserView
	default:
		return nil
	}
}
