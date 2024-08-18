package mapper

import (
	"database/sql"
	"fmt"

	"github.com/jibaru/inheritance-mappers/internal/domain"
)

// BowlerMapper maps Bowler objects to the database.
type BowlerMapper struct {
	CricketerMapper
}

func NewBowlerMapper(db *sql.DB) *BowlerMapper {
	return &BowlerMapper{CricketerMapper: *NewCricketerMapper(db)}
}

func (bm *BowlerMapper) Save(bowler *domain.Bowler) error {
	err := bm.CricketerMapper.Save(&bowler.Cricketer)
	if err != nil {
		return err
	}
	query := "INSERT INTO bowlers (id, bowling_average) VALUES (?, ?)"
	_, err = bm.db.Exec(query, bowler.ID, bowler.BowlingAverage)
	if err != nil {
		return fmt.Errorf("failed to save bowler: %v", err)
	}
	return nil
}

func (bm *BowlerMapper) Load(id int) (*domain.Bowler, error) {
	cricketer, err := bm.CricketerMapper.Load(id)
	if err != nil {
		return nil, err
	}
	var bowler domain.Bowler
	bowler.Cricketer = *cricketer
	query := "SELECT bowling_average FROM bowlers WHERE id = ?"
	err = bm.db.QueryRow(query, id).Scan(&bowler.BowlingAverage)
	if err != nil {
		return nil, fmt.Errorf("failed to load bowler: %v", err)
	}
	return &bowler, nil
}
