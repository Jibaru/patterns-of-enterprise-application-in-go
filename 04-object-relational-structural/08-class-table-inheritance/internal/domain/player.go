package domain

import (
	"database/sql"
	"fmt"
)

// Player represents the base class for all players.
type Player struct {
	ID   int
	Name string
}

// Footballer represents a football player with a club.
type Footballer struct {
	Player
	Club string
}

// Cricketer represents a cricketer with a batting average.
type Cricketer struct {
	Player
	BattingAverage float64
}

// Bowler represents a bowler with a bowling average.
type Bowler struct {
	Cricketer
	BowlingAverage float64
}

// PlayerMapper handles the persistence and retrieval of Player entities.
type PlayerMapper struct {
	DB *sql.DB
}

// NewPlayerMapper creates a new PlayerMapper.
func NewPlayerMapper(db *sql.DB) *PlayerMapper {
	return &PlayerMapper{DB: db}
}

// InsertFootballer inserts a Footballer into the players and footballers tables.
func (mapper *PlayerMapper) InsertFootballer(footballer *Footballer) error {
	// Insert into Players table
	insertPlayerQuery := `INSERT INTO players (name) VALUES (?)`
	result, err := mapper.DB.Exec(insertPlayerQuery, footballer.Name)
	if err != nil {
		return fmt.Errorf("failed to insert player: %w", err)
	}

	playerID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last inserted player id: %w", err)
	}

	// Insert into Footballers table
	insertFootballerQuery := `INSERT INTO footballers (id, club) VALUES (?, ?)`
	_, err = mapper.DB.Exec(insertFootballerQuery, playerID, footballer.Club)
	if err != nil {
		return fmt.Errorf("failed to insert footballer: %w", err)
	}

	footballer.ID = int(playerID)

	return nil
}

// InsertCricketer inserts a Cricketer into the players and cricketers tables.
func (mapper *PlayerMapper) InsertCricketer(cricketer *Cricketer) error {
	// Insert into Players table
	insertPlayerQuery := `INSERT INTO players (name) VALUES (?)`
	result, err := mapper.DB.Exec(insertPlayerQuery, cricketer.Name)
	if err != nil {
		return fmt.Errorf("failed to insert player: %w", err)
	}

	playerID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last inserted player id: %w", err)
	}

	// Insert into Cricketers table
	insertCricketerQuery := `INSERT INTO cricketers (id, batting_average) VALUES (?, ?)`
	_, err = mapper.DB.Exec(insertCricketerQuery, playerID, cricketer.BattingAverage)
	if err != nil {
		return fmt.Errorf("failed to insert cricketer: %w", err)
	}

	cricketer.ID = int(playerID)

	return nil
}

// InsertBowler inserts a Bowler into the players, cricketers, and bowlers tables.
func (mapper *PlayerMapper) InsertBowler(bowler *Bowler) error {
	// Insert into Players table
	insertPlayerQuery := `INSERT INTO players (name) VALUES (?)`
	result, err := mapper.DB.Exec(insertPlayerQuery, bowler.Name)
	if err != nil {
		return fmt.Errorf("failed to insert player: %w", err)
	}

	playerID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last inserted player id: %w", err)
	}

	// Insert into Cricketers table
	insertCricketerQuery := `INSERT INTO cricketers (id, batting_average) VALUES (?, ?)`
	_, err = mapper.DB.Exec(insertCricketerQuery, playerID, bowler.BattingAverage)
	if err != nil {
		return fmt.Errorf("failed to insert cricketer: %w", err)
	}

	// Insert into Bowlers table
	insertBowlerQuery := `INSERT INTO bowlers (id, bowling_average) VALUES (?, ?)`
	_, err = mapper.DB.Exec(insertBowlerQuery, playerID, bowler.BowlingAverage)
	if err != nil {
		return fmt.Errorf("failed to insert bowler: %w", err)
	}

	bowler.ID = int(playerID)

	return nil
}

// GetFootballerByID retrieves a Footballer by ID.
func (mapper *PlayerMapper) GetFootballerByID(id int) (*Footballer, error) {
	query := `
        SELECT p.id, p.name, f.club
        FROM players p
        JOIN footballers f ON p.id = f.id
        WHERE p.id = ?
    `

	row := mapper.DB.QueryRow(query, id)

	var footballer Footballer
	err := row.Scan(&footballer.ID, &footballer.Name, &footballer.Club)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("footballer with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve footballer: %w", err)
	}

	return &footballer, nil
}

// GetCricketerByID retrieves a Cricketer by ID.
func (mapper *PlayerMapper) GetCricketerByID(id int) (*Cricketer, error) {
	query := `
        SELECT p.id, p.name, c.batting_average
        FROM players p
        JOIN cricketers c ON p.id = c.id
        WHERE p.id = ?
    `

	row := mapper.DB.QueryRow(query, id)

	var cricketer Cricketer
	err := row.Scan(&cricketer.ID, &cricketer.Name, &cricketer.BattingAverage)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cricketer with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve cricketer: %w", err)
	}

	return &cricketer, nil
}

// GetBowlerByID retrieves a Bowler by ID.
func (mapper *PlayerMapper) GetBowlerByID(id int) (*Bowler, error) {
	query := `
        SELECT p.id, p.name, c.batting_average, b.bowling_average
        FROM players p
        JOIN cricketers c ON p.id = c.id
        JOIN bowlers b ON c.id = b.id
        WHERE p.id = ?
    `

	row := mapper.DB.QueryRow(query, id)

	var bowler Bowler
	err := row.Scan(&bowler.ID, &bowler.Name, &bowler.BattingAverage, &bowler.BowlingAverage)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("bowler with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve bowler: %w", err)
	}

	return &bowler, nil
}
