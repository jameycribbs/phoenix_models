package phoenix_models

import (
	"time"
)

type ClientApplication struct {
	ID                                         uint      `gorm:"primary_key" json:"id"`
	CreatedAt                                  time.Time `json:"created_at"`
	UpdatedAt                                  time.Time `json:"updated_at"`
	Client                                     Client    `json:"client"`
	ClientID                                   uint      `json:"client_id"`
	Program                                    Program   `json:"program"`
	ProgramID                                  uint      `json:"program_id"`
	CaseManagementActivityDueThisMonthType     string    `json:"case_management_activity_due_this_month"`
	CaseManagementActivityDueThisMonthDateDue  time.Time `json:"case_management_activity_due_this_month_date_due"`
	CaseManagementActivityDueThisMonthDateDone time.Time `json:"case_management_activity_due_this_month_date_done"`
	NextQuarterlyVisitDateDue                  time.Time `json:"next_quarterly_visit_date_due"`
	NextReEvaluationDateDue                    time.Time `json:"next_re_evaluation_date_due"`
}
