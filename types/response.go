package types

type Status string

const (
	SUCCESS Status = "S"
	FAILURE Status = "F"
)

type GenericSuccessResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
}

type ExpensesListResponse struct {
	Status Status        `json:"status"`
	Total  ExpenseAmount `json:"total"`
	//Average  ExpenseAmount `json:"average"`
	Expenses []Expense `json:"expenses"`
}
