package repository

import (
	"database/sql"
	"expense-tracker/structs"
	"strconv"
	"strings"
)

func GetReport(db *sql.DB, userID int, filters map[string]interface{}) (result []structs.Expense, grandTotal float64, err error) {
	// Base query
	query := "SELECT id, user_id, category_id, types, dates, amount, descriptions FROM expenses"

	// Build WHERE clauses
	var conditions []string
	var args []interface{}

	queryGrandTotal := "SELECT COALESCE(SUM(amount), 0) as grandTotal FROM expenses"

	// Always filter by user_id
	conditions = append(conditions, "user_id = $1")
	args = append(args, userID)
	argPos := 2 // Start next parameter at position 2

	if startDate, exists := filters["start_date"]; exists {
		conditions = append(conditions, "dates >= $"+strconv.Itoa(argPos))
		args = append(args, startDate)
		argPos++
	}

	if endDate, exists := filters["end_date"]; exists {
		conditions = append(conditions, "dates <= $"+strconv.Itoa(argPos))
		args = append(args, endDate)
		argPos++
	}

	if itemType, exists := filters["type"]; exists {
		conditions = append(conditions, "types = $"+strconv.Itoa(argPos))
		args = append(args, itemType)
		argPos++
	}

	// Combine conditions if any exist
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
		queryGrandTotal += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Add sorting
	query += " ORDER BY dates DESC"

	// Execute query
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, 0.0, err
	}
	defer rows.Close()

	var expenses []structs.Expense
	for rows.Next() {
		var expense structs.Expense
		err := rows.Scan(&expense.ID, &expense.UserID, &expense.CategoryID, &expense.Types, &expense.Dates, &expense.Amount, &expense.Description)
		if err != nil {
			return nil, 0.0, err
		}
		expenses = append(expenses, expense)
	}

	// execute grand total query
	err = db.QueryRow(queryGrandTotal, args...).Scan(&grandTotal)
	if err != nil {
		return nil, 0.0, err
	}

	return expenses, grandTotal, nil
}
