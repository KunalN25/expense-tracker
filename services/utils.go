package services

import (
	"encoding/json"
	"expense-tracker/types"
	"net/http"
)

func sendJsonResponse(w http.ResponseWriter, jsonData any) {
	w.Header().Set("Content-Type", "application/json")
	// Marshal the Expense struct to JSON
	jsonResponse, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		return

	}
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	successResponse := types.GenericSuccessResponse{
		Status:  types.SUCCESS,
		Message: message,
	}
	sendJsonResponse(w, successResponse)
}

func calculateTotalExpenses(expenses []types.Expense) types.ExpenseAmount {
	var total types.ExpenseAmount
	for _, exp := range expenses {
		total += exp.Amount
	}
	return total
}
