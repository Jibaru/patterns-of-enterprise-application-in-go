package mapper

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/inheritance-mappers/internal/domain"
)

// FootballerMapper maps the Footballer objects to the database.
type FootballerMapper struct {
	PlayerMapper
}

func NewFootballerMapper(db *sql.DB) *FootballerMapper {
	return &FootballerMapper{PlayerMapper: *NewPlayerMapper(db)}
}

func (fm *FootballerMapper) Save(footballer *domain.Footballer) error {
	err := fm.PlayerMapper.Save(&footballer.Player)
	if err != nil {
		return err
	}
	query := "INSERT INTO footballers (id, club) VALUES (?, ?)"
	_, err = fm.db.Exec(query, footballer.ID, footballer.Club)
	if err != nil {
		return fmt.Errorf("failed to save footballer: %v", err)
	}
	return nil
}

func (fm *FootballerMapper) Load(id int) (*domain.Footballer, error) {
	player, err := fm.PlayerMapper.Load(id)
	if err != nil {
		return nil, err
	}
	var footballer domain.Footballer
	footballer.Player = *player
	query := "SELECT club FROM footballers WHERE id = ?"
	err = fm.db.QueryRow(query, id).Scan(&footballer.Club)
	if err != nil {
		return nil, fmt.Errorf("failed to load footballer: %v", err)
	}
	return &footballer, nil
}
