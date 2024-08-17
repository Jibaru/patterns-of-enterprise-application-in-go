package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jibaru/class-table-inheritance/internal/db"
	"github.com/jibaru/class-table-inheritance/internal/domain"
)

func main() {
	// Setup the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("failed to set up database: %v", err)
	}

	playerMapper := domain.NewPlayerMapper(database)

	// Insert a Footballer
	footballer := &domain.Footballer{
		Player: domain.Player{
			Name: "John Doe",
		},
		Club: "FC Barcelona",
	}

	err = playerMapper.InsertFootballer(footballer)
	if err != nil {
		log.Fatalf("failed to insert footballer: %v", err)
	}
	fmt.Println("Footballer inserted successfully")

	// Insert a Cricketer
	cricketer := &domain.Cricketer{
		Player: domain.Player{
			Name: "Sachin Tendulkar",
		},
		BattingAverage: 55.5,
	}

	err = playerMapper.InsertCricketer(cricketer)
	if err != nil {
		log.Fatalf("failed to insert cricketer: %v", err)
	}
	fmt.Println("Cricketer inserted successfully")

	// Insert a Bowler
	bowler := &domain.Bowler{
		Cricketer: domain.Cricketer{
			Player: domain.Player{
				Name: "Shane Warne",
			},
			BattingAverage: 22.4,
		},
		BowlingAverage: 25.2,
	}

	err = playerMapper.InsertBowler(bowler)
	if err != nil {
		log.Fatalf("failed to insert bowler: %v", err)
	}
	fmt.Println("Bowler inserted successfully")

	// Retrieve a Footballer by ID
	retrievedFootballer, err := playerMapper.GetFootballerByID(footballer.ID)
	if err != nil {
		log.Fatalf("failed to retrieve footballer: %v", err)
	}
	data, _ := json.Marshal(retrievedFootballer)
	fmt.Printf("Retrieved Footballer: %v\n", string(data))

	// Retrieve a Cricketer by ID
	retrievedCricketer, err := playerMapper.GetCricketerByID(cricketer.ID)
	if err != nil {
		log.Fatalf("failed to retrieve cricketer: %v", err)
	}
	data, _ = json.Marshal(retrievedCricketer)
	fmt.Printf("Retrieved Cricketer: %v\n", string(data))

	// Retrieve a Bowler by ID
	retrievedBowler, err := playerMapper.GetBowlerByID(bowler.ID)
	if err != nil {
		log.Fatalf("failed to retrieve bowler: %v", err)
	}
	data, _ = json.Marshal(retrievedBowler)
	fmt.Printf("Retrieved Bowler: %v\n", string(data))
}
