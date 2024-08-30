package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jibaru/remote-facade-server/internal/domain"
)

var address = domain.Address{
	Street: "-",
	City:   "-",
	Zip:    "-",
}

func GetAddressHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /address")
	json.NewEncoder(w).Encode(address)
	log.Printf("retrieving address: %v\n", address)
}

func SetAddressHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /address")

	decoder := json.NewDecoder(r.Body)
	var addr domain.Address
	if err := decoder.Decode(&addr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	address = addr

	log.Printf("new address: %v\n", address)

	w.WriteHeader(http.StatusOK)
}
