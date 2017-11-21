package phoenix_models

import (
	"time"
)

type EligibilityCategory struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Value     string    `json:"value"`
}
