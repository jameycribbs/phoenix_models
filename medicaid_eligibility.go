package phoenix_models

import (
	"time"
)

type MedicaidEligibility struct {
	ID                  uint                `gorm:"primary_key" json:"id"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
	EligibilityCategory EligibilityCategory `gorm:"ForeignKey:PaymentCategory" json:"eligibility_category"`
	PaymentCategory     string              `json:"payment_category"`
}
