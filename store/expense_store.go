package store

import (
	"errors"
	"expense-tracker/types"
)

func GetExpenses() ([]types.Expense, error) {
	expenses, err := LoadData()
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func GetExpenseById(id string) (types.Expense, error) {
	expenses, err := LoadData()
	if err != nil {
		return types.Expense{}, err
	}
	for _, exp := range expenses {
		if exp.Id == id {
			return exp, nil
		}
	}

	return types.Expense{}, nil
}

func Save(expense types.Expense) error {
	expenses, err := LoadData()
	if err != nil {
		return err
	}
	expenseExists := false
	for idx, exp := range expenses {
		if exp.Id == expense.Id {
			expenseExists = true
			expenses[idx] = expense
		}
	}
	if !expenseExists {
		expenses = append(expenses, expense)
	}

	err = SaveData(expenses)
	if err != nil {
		return err
	}
	return nil
}

func SaveAll(newExpenses []types.Expense) error {
	expensesData, err := LoadData()
	if err != nil {
		return err
	}
	expensesData = append(expensesData, newExpenses...)

	err = SaveData(expensesData)
	if err != nil {
		return err
	}
	return nil
}

func Delete(expenseId string) error {
	expenses, err := LoadData()
	if err != nil {
		return err
	}
	idx := -1
	for i, exp := range expenses {
		if exp.Id == expenseId {
			idx = i
			break
		}
	}
	if idx == -1 {
		return errors.New("Expense not found with id " + expenseId)
	}
	expenses = append(expenses[:idx], expenses[idx+1:]...)
	err = SaveData(expenses)
	return err
}
