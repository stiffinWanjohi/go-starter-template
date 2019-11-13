package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/etowett/datsimple/backend/data"
	"github.com/etowett/datsimple/backend/db"
	"github.com/etowett/datsimple/backend/repos"
)

type (
	DailyCostService interface {
		AddDailyCost(*db.DBManager, *data.CreateDailyCostForm) (*data.DailyCost, error)
	}

	SimpleDailyCostService struct {
		dailyCostRepository repos.DailyCostRepository
	}
)

func NewDailyCostService(
	dailyCostRepository repos.DailyCostRepository,
) *SimpleDailyCostService {
	return &SimpleDailyCostService{
		dailyCostRepository: dailyCostRepository,
	}
}

func (s *SimpleDailyCostService) AddDailyCost(
	dbManager *db.DBManager,
	costForm *data.CreateDailyCostForm,
) (*data.DailyCost, error) {
	targetDate, err := time.Parse("2006-01-02", costForm.TargetDate)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error with given date [%v], expected date in format [yyyy-mm-dd]", costForm.TargetDate))
	}
	return s.dailyCostRepository.InsertDailyCost(dbManager, &data.DailyCost{
		UserID:     costForm.UserID,
		TargetDate: targetDate,
		Amount:     costForm.Amount,
		CreatedAt:  time.Now(),
	})
}
