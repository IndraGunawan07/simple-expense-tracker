package repository

import (
	"database/sql"
	"expense-tracker/structs"
	"fmt"
	"time"
)

func GetAllExpense(db *sql.DB, userID int) (result []structs.Expense, err error) {
	sql := "SELECT * FROM expenses WHERE user_id = $1"

	rows, err := db.Query(sql, userID)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var expense structs.Expense

		err = rows.Scan(&expense.ID, &expense.UserID, &expense.CategoryID, &expense.Types, &expense.Dates, &expense.CreatedAt, &expense.UpdatedAt)
		if err != nil {
			return
		}

		result = append(result, expense)
	}

	return
}

func InsertExpense(db *sql.DB, expense structs.Expense) (int, error) {
	sql := "INSERT INTO expenses(user_id, category_id, types, dates, amount, descriptions, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"

	var id int

	now := time.Now()

	errs := db.QueryRow(sql, expense.UserID, expense.CategoryID, expense.Types, expense.Dates, expense.Amount, expense.Description, now, now).Scan(&id)
	if errs != nil {
		return 0, fmt.Errorf("failed to insert category: %w", errs)
	}

	return id, nil
}

func UpdateExpense(db *sql.DB, expense structs.Expense) (err error) {
	sql := "UPDATE expenses SET category_id = $1, types = $2, dates = $3, amount = $4, descriptions = $5 updated_at = $6 WHERE id = $7"

	errs := db.QueryRow(sql, expense.CategoryID, expense.Types, expense.Dates, expense.Amount, expense.Description, time.Now(), expense.ID)

	return errs.Err()
}

func DeleteExpense(db *sql.DB, expense structs.Expense) (err error) {
	sql := "DELETE FROM expenses WHERE id = $1"

	errs := db.QueryRow(sql, expense.ID)
	return errs.Err()
}
