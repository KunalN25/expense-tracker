package routes

import (
	"expense-tracker/services"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/getExpenses", services.GetExpenses)
	mux.HandleFunc("/addExpense", services.AddExpense)
	mux.HandleFunc("/editExpense", services.EditExpense)
	mux.HandleFunc("/deleteExpense", services.DeleteExpense)
}
