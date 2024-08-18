package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jibaru/inheritance-mappers/internal/db"
	"github.com/jibaru/inheritance-mappers/internal/domain"
	"github.com/jibaru/inheritance-mappers/internal/mapper"
)

func main() {
	// Set up the database connection
	database, err := db.Setup()
	if err != nil {
		log.Fatalf("Could not set up the database: %v", err)
	}
	defer database.Close()

	// Demonstrating Footballer insertion and retrieval
	footballerMapper := mapper.NewFootballerMapper(database)
	footballer := &domain.Footballer{
		Player: domain.Player{
			ID:   1,
			Name: "John Doe",
		},
		Club: "FC Barcelona",
	}
	err = footballerMapper.Save(footballer)
	if err != nil {
		log.Fatalf("Could not save footballer: %v", err)
	}

	retrievedFootballer, err := footballerMapper.Load(1)
	if err != nil {
		log.Fatalf("Could not load footballer: %v", err)
	}
	data, _ := json.Marshal(retrievedFootballer)
	fmt.Printf("Retrieved Footballer: %v\n", string(data))

	// Demonstrating Cricketer insertion and retrieval
	cricketerMapper := mapper.NewCricketerMapper(database)
	cricketer := &domain.Cricketer{
		Player: domain.Player{
			ID:   2,
			Name: "James Smith",
		},
		BattingAverage: 45.3,
	}
	err = cricketerMapper.Save(cricketer)
	if err != nil {
		log.Fatalf("Could not save cricketer: %v", err)
	}

	retrievedCricketer, err := cricketerMapper.Load(2)
	if err != nil {
		log.Fatalf("Could not load cricketer: %v", err)
	}
	data, _ = json.Marshal(retrievedCricketer)
	fmt.Printf("Retrieved Cricketer: %v\n", string(data))

	// Demonstrating Bowler insertion and retrieval
	bowlerMapper := mapper.NewBowlerMapper(database)
	bowler := &domain.Bowler{
		Cricketer: domain.Cricketer{
			Player: domain.Player{
				ID:   3,
				Name: "David Miller",
			},
			BattingAverage: 35.4,
		},
		BowlingAverage: 25.6,
	}
	err = bowlerMapper.Save(bowler)
	if err != nil {
		log.Fatalf("Could not save bowler: %v", err)
	}

	retrievedBowler, err := bowlerMapper.Load(3)
	if err != nil {
		log.Fatalf("Could not load bowler: %v", err)
	}
	data, _ = json.Marshal(retrievedBowler)
	fmt.Printf("Retrieved Bowler: %v\n", string(data))
}
