package types

import "time"

type AddExpenseRequest struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Amount      ExpenseAmount `json:"amount"`
}

type EditExpenseRequest struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
}

type DeleteExpenseRequest struct {
	Id string `json:"id"`
}

type GetExpensesRequest struct {
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty"`
}
