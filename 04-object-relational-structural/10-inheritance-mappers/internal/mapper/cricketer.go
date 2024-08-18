package mapper

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/inheritance-mappers/internal/domain"
)

// CricketerMapper maps Cricketer objects to the database.
type CricketerMapper struct {
	PlayerMapper
}

func NewCricketerMapper(db *sql.DB) *CricketerMapper {
	return &CricketerMapper{PlayerMapper: *NewPlayerMapper(db)}
}

func (cm *CricketerMapper) Save(cricketer *domain.Cricketer) error {
	err := cm.PlayerMapper.Save(&cricketer.Player)
	if err != nil {
		return err
	}
	query := "INSERT INTO cricketers (id, batting_average) VALUES (?, ?)"
	_, err = cm.db.Exec(query, cricketer.ID, cricketer.BattingAverage)
	if err != nil {
		return fmt.Errorf("failed to save cricketer: %v", err)
	}
	return nil
}

func (cm *CricketerMapper) Load(id int) (*domain.Cricketer, error) {
	player, err := cm.PlayerMapper.Load(id)
	if err != nil {
		return nil, err
	}
	var cricketer domain.Cricketer
	cricketer.Player = *player
	query := "SELECT batting_average FROM cricketers WHERE id = ?"
	err = cm.db.QueryRow(query, id).Scan(&cricketer.BattingAverage)
	if err != nil {
		return nil, fmt.Errorf("failed to load cricketer: %v", err)
	}
	return &cricketer, nil
}
