package routes

import (
	"github.com/gorilla/mux"
	"loanengine.com/mod/handlers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/loans", handlers.CreateLoan).Methods("POST")
	r.HandleFunc("/loans/{id}/approve", handlers.ApproveLoan).Methods("POST")
	r.HandleFunc("/loans/{id}/invest", handlers.InvestLoan).Methods("POST")
	r.HandleFunc("/loans/{id}/disburse", handlers.DisburseLoan).Methods("POST")
	r.HandleFunc("/loans/{id}", handlers.GetLoan).Methods("GET")

	return r
}
