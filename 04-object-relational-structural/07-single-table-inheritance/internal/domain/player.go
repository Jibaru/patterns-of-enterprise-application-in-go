package domain

import (
	"database/sql"
	"fmt"
)

// Player represents a base player entity.
type Player struct {
	ID   int
	Name string
	Type string
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

// PlayerMapper handles persistence and retrieval of Player entities.
type PlayerMapper struct {
	DB *sql.DB
}

// NewPlayerMapper creates a new PlayerMapper.
func NewPlayerMapper(db *sql.DB) *PlayerMapper {
	return &PlayerMapper{DB: db}
}

// InsertFootballer inserts a Footballer into the database.
func (mapper *PlayerMapper) InsertFootballer(footballer *Footballer) error {
	query := `INSERT INTO players (name, club, type) VALUES (?, ?, 'footballer')`
	result, err := mapper.DB.Exec(query, footballer.Name, footballer.Club)
	if err != nil {
		return fmt.Errorf("failed to insert footballer: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to insert bowler: %w", err)
	}

	footballer.ID = int(id)

	return nil
}

// InsertCricketer inserts a Cricketer into the database.
func (mapper *PlayerMapper) InsertCricketer(cricketer *Cricketer) error {
	query := `INSERT INTO players (name, batting_average, type) VALUES (?, ?, 'cricketer')`
	result, err := mapper.DB.Exec(query, cricketer.Name, cricketer.BattingAverage)
	if err != nil {
		return fmt.Errorf("failed to insert cricketer: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to insert bowler: %w", err)
	}

	cricketer.ID = int(id)

	return nil
}

// InsertBowler inserts a Bowler into the database.
func (mapper *PlayerMapper) InsertBowler(bowler *Bowler) error {
	query := `INSERT INTO players (name, batting_average, bowling_average, type) VALUES (?, ?, ?, 'bowler')`
	result, err := mapper.DB.Exec(query, bowler.Name, bowler.BattingAverage, bowler.BowlingAverage)
	if err != nil {
		return fmt.Errorf("failed to insert bowler: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to insert bowler: %w", err)
	}

	bowler.ID = int(id)

	return nil
}

// GetFootballer retrieve a Footballer for its id
func (mapper *PlayerMapper) GetFootballer(id int) (*Footballer, error) {
	query := `SELECT id, name, type, club FROM players WHERE id = ? and type = ?`
	row := mapper.DB.QueryRow(query, id, "footballer")

	var f Footballer
	err := row.Scan(&f.ID, &f.Name, &f.Type, &f.Club)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("footballer with id %v not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve footballer: %w", err)
	}

	return &f, nil
}

// GetCricketer retrieve a Cricketer for its id
func (mapper *PlayerMapper) GetCricketer(id int) (*Cricketer, error) {
	query := `SELECT id, name, type, batting_average FROM players WHERE id = ? and type = ?`
	row := mapper.DB.QueryRow(query, id, "cricketer")

	var f Cricketer
	err := row.Scan(&f.ID, &f.Name, &f.Type, &f.BattingAverage)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cricketer with id %v not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve cricketer: %w", err)
	}

	return &f, nil
}

// GetBowler retrieve a Bowler for its id
func (mapper *PlayerMapper) GetBowler(id int) (*Bowler, error) {
	query := `SELECT id, name, type, batting_average, bowling_average FROM players WHERE id = ? and type = ?`
	row := mapper.DB.QueryRow(query, id, "bowler")

	var f Bowler
	err := row.Scan(&f.ID, &f.Name, &f.Type, &f.BattingAverage, &f.BowlingAverage)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("bowler with id %v not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve bowler: %w", err)
	}

	return &f, nil
}
