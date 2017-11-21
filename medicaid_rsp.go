package phoenix_models

import (
	"time"
)

type MedicaidRsp struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	MedicaidNumber string    `json:"medicaid_number"`
}
