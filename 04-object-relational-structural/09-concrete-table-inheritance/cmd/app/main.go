package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jibaru/concrete-table-inheritance/internal/db"
	"github.com/jibaru/concrete-table-inheritance/internal/domain"
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
			Name: "Lionel Messi",
		},
		Club: "Paris Saint-Germain",
	}

	err = playerMapper.InsertFootballer(footballer)
	if err != nil {
		log.Fatalf("failed to insert footballer: %v", err)
	}
	fmt.Printf("Footballer inserted with ID %d\n", footballer.ID)

	// Insert a Cricketer
	cricketer := &domain.Cricketer{
		Player: domain.Player{
			Name: "Virat Kohli",
		},
		BattingAverage: 58.2,
	}

	err = playerMapper.InsertCricketer(cricketer)
	if err != nil {
		log.Fatalf("failed to insert cricketer: %v", err)
	}
	fmt.Printf("Cricketer inserted with ID %d\n", cricketer.ID)

	// Insert a Bowler
	bowler := &domain.Bowler{
		Cricketer: domain.Cricketer{
			Player: domain.Player{
				Name: "Mitchell Starc",
			},
			BattingAverage: 15.3,
		},
		BowlingAverage: 24.1,
	}

	err = playerMapper.InsertBowler(bowler)
	if err != nil {
		log.Fatalf("failed to insert bowler: %v", err)
	}
	fmt.Printf("Bowler inserted with ID %d\n", bowler.ID)

	// Retrieve and display a Footballer by ID
	retrievedFootballer, err := playerMapper.GetFootballerByID(footballer.ID)
	if err != nil {
		log.Fatalf("failed to retrieve footballer: %v", err)
	}
	data, _ := json.Marshal(retrievedFootballer)
	fmt.Printf("Retrieved Footballer: %v\n", string(data))

	// Retrieve and display a Cricketer by ID
	retrievedCricketer, err := playerMapper.GetCricketerByID(cricketer.ID)
	if err != nil {
		log.Fatalf("failed to retrieve cricketer: %v", err)
	}
	data, _ = json.Marshal(retrievedCricketer)
	fmt.Printf("Retrieved Cricketer: %v\n", string(data))

	// Retrieve and display a Bowler by ID
	retrievedBowler, err := playerMapper.GetBowlerByID(bowler.ID)
	if err != nil {
		log.Fatalf("failed to retrieve bowler: %v", err)
	}
	data, _ = json.Marshal(retrievedBowler)
	fmt.Printf("Retrieved Bowler: %v\n", string(data))
}
