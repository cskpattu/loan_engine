package models

import "time"

type LoanState string

const (
	StateProposed  LoanState = "proposed"
	StateApproved  LoanState = "approved"
	StateInvested  LoanState = "invested"
	StateDisbursed LoanState = "disbursed"
)

type ApprovalInfo struct {
	ProofPictureURL string    `json:"proof_picture_url"`
	ValidatorID     string    `json:"validator_id"`
	Date            time.Time `json:"date"`
}

type Investment struct {
	InvestorID string  `json:"investor_id"`
	Amount     float64 `json:"amount"`
}

type DisbursementInfo struct {
	AgreementLetterURL string    `json:"agreement_letter_url"`
	OfficerID          string    `json:"officer_id"`
	Date               time.Time `json:"date"`
}

type Loan struct {
	ID                  string            `json:"id"`
	BorrowerID          string            `json:"borrower_id"`
	PrincipalAmount     float64           `json:"principal_amount"`
	Rate                float64           `json:"rate"`
	ROI                 float64           `json:"roi"`
	AgreementLetterURL  string            `json:"agreement_letter_url"`
	State               LoanState         `json:"state"`
	Approval            *ApprovalInfo     `json:"approval,omitempty"`
	Investments         []Investment      `json:"investments,omitempty"`
	Disbursement        *DisbursementInfo `json:"disbursement,omitempty"`
	TotalInvestedAmount float64           `json:"total_invested_amount"`
}
