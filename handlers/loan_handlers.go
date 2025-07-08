package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"loanengine.com/mod/models"
	"loanengine.com/mod/utils"
	"net/http"
	"sync"
	"time"
)

var (
	loans   = make(map[string]*models.Loan)
	loanMux = sync.RWMutex{}
)

func CreateLoan(w http.ResponseWriter, r *http.Request) {
	var loan models.Loan
	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loanMux.Lock()
	defer loanMux.Unlock()

	loan.ID = fmt.Sprintf("loan-%d", len(loans)+1)
	loan.State = models.StateProposed
	loans[loan.ID] = &loan

	json.NewEncoder(w).Encode(loan)
}

func ApproveLoan(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var approval models.ApprovalInfo

	if err := json.NewDecoder(r.Body).Decode(&approval); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loanMux.Lock()
	defer loanMux.Unlock()

	loan, exists := loans[id]
	if !exists || loan.State != models.StateProposed {
		http.Error(w, "Loan not found or not in proposed state", http.StatusBadRequest)
		return
	}

	approval.Date = time.Now()
	loan.Approval = &approval
	loan.State = models.StateApproved

	json.NewEncoder(w).Encode(loan)
}

func InvestLoan(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var inv models.Investment

	if err := json.NewDecoder(r.Body).Decode(&inv); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loanMux.Lock()
	defer loanMux.Unlock()

	loan, exists := loans[id]
	if !exists || (loan.State != models.StateApproved && loan.State != models.StateInvested) {
		http.Error(w, "Loan not ready for investment", http.StatusBadRequest)
		return
	}

	if loan.TotalInvestedAmount+inv.Amount > loan.PrincipalAmount {
		http.Error(w, "Investment exceeds principal", http.StatusBadRequest)
		return
	}

	loan.Investments = append(loan.Investments, inv)
	loan.TotalInvestedAmount += inv.Amount

	if loan.TotalInvestedAmount == loan.PrincipalAmount {
		loan.State = models.StateInvested
		utils.SendAgreementEmails(loan)
	}

	json.NewEncoder(w).Encode(loan)
}

func DisburseLoan(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var disb models.DisbursementInfo

	if err := json.NewDecoder(r.Body).Decode(&disb); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loanMux.Lock()
	defer loanMux.Unlock()

	loan, exists := loans[id]
	if !exists || loan.State != models.StateInvested {
		http.Error(w, "Loan not ready for disbursement", http.StatusBadRequest)
		return
	}

	disb.Date = time.Now()
	loan.Disbursement = &disb
	loan.State = models.StateDisbursed

	json.NewEncoder(w).Encode(loan)
}

func GetLoan(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	loanMux.RLock()
	defer loanMux.RUnlock()

	loan, exists := loans[id]
	if !exists {
		http.Error(w, "Loan not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(loan)
}
