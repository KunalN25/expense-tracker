package store

import (
	"encoding/json"
	"expense-tracker/types"
	"os"
)

const filename = "store/data/expensesData.json"

func LoadData() ([]types.Expense, error) {
	var expenses []types.Expense
	file, err := os.Open(filename)
	if err == nil {
		// File exists, so read the existing todos
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&expenses); err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	return expenses, nil
}

func SaveData(expenses []types.Expense) error {
	file, err := os.Create(filename)
	if err != nil {
		return err // Return error if file creation fails
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: for pretty-printing
	if err := encoder.Encode(expenses); err != nil {
		return err // Return error if encoding fails
	}

	return nil // Return nil if successful
}
