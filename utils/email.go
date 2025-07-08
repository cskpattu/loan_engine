package utils

import (
	"fmt"
	"loanengine.com/mod/models"
)

func SendAgreementEmails(loan *models.Loan) {
	for _, investor := range loan.Investments {
		fmt.Printf("Email sent to investor %s for loan %s\n", investor.InvestorID, loan.ID)
	}
}
