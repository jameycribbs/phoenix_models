package phoenix_models

import (
	"time"
)

type MedicaidEligibility struct {
	ID                  uint                `gorm:"primary_key" json:"id"`
	CreatedAt           time.Time           `json:"created-at"`
	UpdatedAt           time.Time           `json:"updated-at"`
	EligibilityCategory EligibilityCategory `gorm:"ForeignKey:PaymentCategory" json:"updated-at"`
	PaymentCategory     string              `json:"payment-category"`
}
