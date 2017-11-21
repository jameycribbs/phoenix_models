package phoenix_models

import (
	"time"
)

type Client struct {
	ID                    uint                `gorm:"primary_key" json:"id"`
	CreatedAt             time.Time           `json:"created_at"`
	UpdatedAt             time.Time           `json:"updated_at"`
	Area                  Area                `json:"area"`
	AreaID                uint                `json:"area_id"`
	MedicaidEligibility   MedicaidEligibility `json:"medicaid_eligibility"`
	MedicaidEligibilityID uint                `json:"medicaid_eligibility_id"`
	MedicaidRsps          []MedicaidRsp       `gorm:"ForeignKey:MedicaidNumber;AssociationForeignKey:MedicaidNumber" json:"medicaid_rsps"`
	FirstName             string              `json:"first_name"`
	LastName              string              `json:"last_name"`
	MedicaidNumber        string              `json:"medicaid_number"`
}
