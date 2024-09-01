package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jibaru/record-set/internal/data"
	"github.com/jibaru/record-set/internal/db"
	"github.com/jibaru/record-set/internal/repositories"
)

func main() {
	// Set up the database connection
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}
	defer dbConn.Close()

	// Initialize the repository and service layers
	userRepo := repositories.NewUserRepository(dbConn)

	rows, err := userRepo.GetAll()
	if err != nil {
		log.Fatalf("Failed to get all users: %v", err)
		return
	}

	for rows.Next() {
		var user data.User
		err = rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName)
		if err != nil {
			log.Fatalf("Failed to scan users: %v", err)
			return
		}

		userJson, _ := json.Marshal(user)

		fmt.Println(string(userJson))
	}
}
