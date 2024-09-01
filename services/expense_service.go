package services

import (
	"encoding/json"
	"expense-tracker/constants"
	"expense-tracker/store"
	"expense-tracker/types"
	"fmt"
	"net/http"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	var getExpensesReq types.GetExpensesRequest
	err := json.NewDecoder(r.Body).Decode(&getExpensesReq)
	fmt.Println(getExpensesReq)
	startDate, endDate := getExpensesReq.StartDate, getExpensesReq.EndDate

	expenses, err := store.GetExpenses()
	if err != nil {
		http.Error(w, "Could not fetch expenses", http.StatusInternalServerError)
		return
	}
	if startDate == nil && endDate == nil {
		expensesListResponse := types.ExpensesListResponse{
			Status:   types.SUCCESS,
			Expenses: expenses,
			Total:    calculateTotalExpenses(expenses),
		}
		sendJsonResponse(w, expensesListResponse)
		return
	}

	var filteredExpenses []types.Expense
	for _, exp := range expenses {
		if (startDate == nil || !exp.CreatedDate.Before(*startDate)) &&
			(endDate == nil || !exp.CreatedDate.After(*endDate)) {
			filteredExpenses = append(filteredExpenses, exp)
		}
	}

	expensesListResponse := types.ExpensesListResponse{
		Status:   types.SUCCESS,
		Expenses: filteredExpenses,
		Total:    calculateTotalExpenses(filteredExpenses),
	}
	sendJsonResponse(w, expensesListResponse)
}

func AddExpense(w http.ResponseWriter, r *http.Request) {
	var addExpenseRequest types.AddExpenseRequest

	err := json.NewDecoder(r.Body).Decode(&addExpenseRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if addExpenseRequest.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if len(addExpenseRequest.Title) > constants.TitleLimit {
		http.Error(w, "Title is too long", http.StatusBadRequest)
		return
	}
	if addExpenseRequest.Amount <= 0.0 {
		http.Error(w, "Please enter valid amount", http.StatusBadRequest)
		return
	}

	expense := types.CreateExpense(addExpenseRequest.Title, addExpenseRequest.Description, addExpenseRequest.Amount)

	err = store.Save(expense)
	if err != nil {
		http.Error(w, "Unable to add expense", http.StatusInternalServerError)
		return
	}

	sendSuccessResponse(w, "Successfully added expense")
}

func EditExpense(w http.ResponseWriter, r *http.Request) {
	var editExpenseReq types.Expense
	err := json.NewDecoder(r.Body).Decode(&editExpenseReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := editExpenseReq.Id
	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	if len(editExpenseReq.Title) > constants.TitleLimit {
		http.Error(w, "Title is too long", http.StatusBadRequest)
		return
	}
	if editExpenseReq.Amount < 0 {
		http.Error(w, "Please enter valid amount", http.StatusBadRequest)
		return
	}
	editExpenseReq.UpdateLastModified()
	err = store.Save(editExpenseReq)
	if err != nil {
		http.Error(w, "Unable to update expense", http.StatusInternalServerError)
		return
	}
	sendSuccessResponse(w, "Successfully updated expense")
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var deleteExpenseReq types.DeleteExpenseRequest
	err := json.NewDecoder(r.Body).Decode(&deleteExpenseReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := deleteExpenseReq.Id
	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}
	err = store.Delete(id)
	if err != nil {
		http.Error(w, "Unable to delete expense", http.StatusInternalServerError)
		return
	}
	sendSuccessResponse(w, "Successfully deleted expense")
}
