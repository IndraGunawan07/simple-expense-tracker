package repository

import (
	"database/sql"
	"expense-tracker/structs"
	"fmt"
)

func GetAllCategory(db *sql.DB) (result []structs.Category, err error) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var category structs.Category

		err = rows.Scan(&category.ID, &category.Nama)
		if err != nil {
			return
		}

		result = append(result, category)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (int, error) {
	sql := "INSERT INTO Category(name) VALUES ($1) RETURNING id"

	var id int

	errs := db.QueryRow(sql, category.Nama).Scan(&id)
	if errs != nil {
		return 0, fmt.Errorf("failed to insert category: %w", errs)
	}

	return id, nil
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1 WHERE id = $2"

	errs := db.QueryRow(sql, category.Nama, category.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)
	return errs.Err()
}
