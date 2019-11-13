package repos

import (
	"github.com/etowett/datsimple/backend/data"
	"github.com/etowett/datsimple/backend/db"
)

type (
	DailyCostRepository interface {
		InsertDailyCost(db.SQLOperations, *data.DailyCost) (*data.DailyCost, error)
	}

	SimpleDailyCostRepository struct{}
)

func NewDailyCostRepository() *SimpleDailyCostRepository {
	return &SimpleDailyCostRepository{}
}

func (r *SimpleDailyCostRepository) InsertDailyCost(
	operations db.SQLOperations,
	cost *data.DailyCost,
) (*data.DailyCost, error) {
	err := operations.QueryRow(
		`insert into daily_cost(user_id, target_date, amount, created_at) VALUES($1, $2, $3, $4) returning id`, cost.UserID, cost.TargetDate, cost.Amount, cost.CreatedAt,
	).Scan(&cost.ID)
	return cost, err
}
