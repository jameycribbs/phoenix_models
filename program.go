package phoenix_models

import (
	"time"
)

type Program struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
	Name      string    `json:"name"`
}
