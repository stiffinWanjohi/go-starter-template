package data

import "time"

type (
	DailyCost struct {
		ID         int64     `json:"id"`
		UserID     int64     `json:"user_id"`
		TargetDate time.Time `json:"target_date"`
		Amount     float64   `json:"amount"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	CreateDailyCostForm struct {
		UserID     int64   `json:"user_id" binding:"required"`
		TargetDate string  `json:"target_date" binding:"required"`
		Amount     float64 `json:"amount" binding:"required"`
	}
)
