package phoenix_models

import (
	"time"
)

type MedicaidRsp struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time `json:"created-at"`
	UpdatedAt      time.Time `json:"updated-at"`
	MedicaidNumber string    `json:"medicaid-number"`
}
