package types

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type ExpenseAmount float32

type Expense struct {
	Id           string        `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Amount       ExpenseAmount `json:"amount"`
	CreatedDate  time.Time     `json:"createdDate"`
	LastModified time.Time     `json:"lastModified"`
}

func CreateExpense(title, description string, amount ExpenseAmount) Expense {

	return Expense{
		Id:           uuid.NewString(),
		Title:        title,
		Description:  description,
		Amount:       amount,
		CreatedDate:  time.Now(),
		LastModified: time.Now(),
	}
}

func (e *Expense) UpdateTitle(title string) {
	e.Title = title
}
func (e *Expense) UpdateAmount(newAmount ExpenseAmount) {
	e.Amount = newAmount
}

func (e *Expense) UpdateDescription(newDescription string) {
	e.Description = newDescription
}

func (e *Expense) UpdateLastModified() {
	e.LastModified = time.Now()
}

func (e *Expense) String() string {
	return fmt.Sprintf("ID: %s\nTitle: %s\nDescription: %s\nAmount: %.2f\nDate: %s",
		e.Id, e.Title, e.Description, e.Amount, e.CreatedDate.Format(time.RFC3339))
}
