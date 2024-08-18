package mapper

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/inheritance-mappers/internal/domain"
)

// AbstractPlayerMapper defines the common functionality for all player mappers.
type AbstractPlayerMapper interface {
	Save(player domain.Player) error
	Load(id int) (domain.Player, error)
}

// PlayerMapper handles basic player mapping.
type PlayerMapper struct {
	db *sql.DB
}

func NewPlayerMapper(db *sql.DB) *PlayerMapper {
	return &PlayerMapper{db: db}
}

func (pm *PlayerMapper) Save(player *domain.Player) error {
	query := "INSERT INTO players (id, name) VALUES (?, ?)"
	_, err := pm.db.Exec(query, player.ID, player.Name)
	if err != nil {
		return fmt.Errorf("failed to save player: %v", err)
	}
	return nil
}

func (pm *PlayerMapper) Load(id int) (*domain.Player, error) {
	var player domain.Player
	query := "SELECT id, name FROM players WHERE id = ?"
	row := pm.db.QueryRow(query, id)
	err := row.Scan(&player.ID, &player.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to load player: %v", err)
	}
	return &player, nil
}
