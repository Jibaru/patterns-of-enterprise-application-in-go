package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jibaru/single-table-inheritance/internal/db"
	"github.com/jibaru/single-table-inheritance/internal/domain"
)

func main() {
	// Initialize the database
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("failed to set up database: %v", err)
	}

	// Create PlayerMapper
	playerMapper := domain.NewPlayerMapper(database)

	// Insert a Footballer
	footballer := &domain.Footballer{
		Player: domain.Player{Name: "Lionel Messi", Type: "footballer"},
		Club:   "Paris Saint-Germain",
	}

	err = playerMapper.InsertFootballer(footballer)
	if err != nil {
		log.Fatalf("failed to insert footballer: %v\n", err)
	}

	fmt.Println("Footballer inserted successfully")

	// Insert a Cricketer
	cricketer := &domain.Cricketer{
		Player:         domain.Player{Name: "Virat Kohli", Type: "cricketer"},
		BattingAverage: 50.5,
	}

	err = playerMapper.InsertCricketer(cricketer)
	if err != nil {
		log.Fatalf("failed to insert cricketer: %v\n", err)
	}

	fmt.Println("Cricketer inserted successfully")

	// Insert a Bowler
	bowler := &domain.Bowler{
		Cricketer: domain.Cricketer{
			Player:         domain.Player{Name: "Shane Warne", Type: "bowler"},
			BattingAverage: 15.8,
		},
		BowlingAverage: 25.4,
	}

	err = playerMapper.InsertBowler(bowler)
	if err != nil {
		log.Fatalf("failed to insert bowler: %v\n", err)
	}

	fmt.Println("Bowler inserted successfully")

	p1, err := playerMapper.GetFootballer(footballer.ID)
	if err != nil {
		log.Fatalf("failed to retrieve player: %v\n", err)
	}
	data, _ := json.Marshal(p1)

	fmt.Printf("Retrieved footballer: %v\n", string(data))

	p2, err := playerMapper.GetCricketer(cricketer.ID)
	if err != nil {
		log.Fatalf("failed to retrieve player: %v\n", err)
	}
	data, _ = json.Marshal(p2)

	fmt.Printf("Retrieved cricketer: %v\n", string(data))

	p3, err := playerMapper.GetBowler(bowler.ID)
	if err != nil {
		log.Fatalf("failed to retrieve player: %v\n", err)
	}
	data, _ = json.Marshal(p3)

	fmt.Printf("Retrieved bowler: %v\n", string(data))
}
