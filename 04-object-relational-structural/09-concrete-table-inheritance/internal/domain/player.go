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

// PlayerMapper handles the persistence and retrieval of Player entities (Footballer, Cricketer, Bowler).
type PlayerMapper struct {
	DB *sql.DB
}

// NewPlayerMapper creates a new PlayerMapper.
func NewPlayerMapper(db *sql.DB) *PlayerMapper {
	return &PlayerMapper{DB: db}
}

// InsertFootballer inserts a Footballer into the footballers table.
func (mapper *PlayerMapper) InsertFootballer(footballer *Footballer) error {
	insertFootballerQuery := `INSERT INTO footballers (name, club) VALUES (?, ?)`
	result, err := mapper.DB.Exec(insertFootballerQuery, footballer.Name, footballer.Club)
	if err != nil {
		return fmt.Errorf("failed to insert footballer: %w", err)
	}

	footballerID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last inserted footballer id: %w", err)
	}
	footballer.ID = int(footballerID)
	return nil
}

// InsertCricketer inserts a Cricketer into the cricketers table.
func (mapper *PlayerMapper) InsertCricketer(cricketer *Cricketer) error {
	insertCricketerQuery := `INSERT INTO cricketers (name, batting_average) VALUES (?, ?)`
	result, err := mapper.DB.Exec(insertCricketerQuery, cricketer.Name, cricketer.BattingAverage)
	if err != nil {
		return fmt.Errorf("failed to insert cricketer: %w", err)
	}

	cricketerID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last inserted cricketer id: %w", err)
	}
	cricketer.ID = int(cricketerID)
	return nil
}

// InsertBowler inserts a Bowler into the bowlers table.
func (mapper *PlayerMapper) InsertBowler(bowler *Bowler) error {
	insertBowlerQuery := `INSERT INTO bowlers (name, batting_average, bowling_average) VALUES (?, ?, ?)`
	result, err := mapper.DB.Exec(insertBowlerQuery, bowler.Name, bowler.BattingAverage, bowler.BowlingAverage)
	if err != nil {
		return fmt.Errorf("failed to insert bowler: %w", err)
	}

	bowlerID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last inserted bowler id: %w", err)
	}
	bowler.ID = int(bowlerID)
	return nil
}

// GetFootballerByID retrieves a Footballer by ID.
func (mapper *PlayerMapper) GetFootballerByID(id int) (*Footballer, error) {
	query := `SELECT id, name, club FROM footballers WHERE id = ?`
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
	query := `SELECT id, name, batting_average FROM cricketers WHERE id = ?`
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
	query := `SELECT id, name, batting_average, bowling_average FROM bowlers WHERE id = ?`
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
