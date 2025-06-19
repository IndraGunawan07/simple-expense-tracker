package repository

import (
	// "expense-tracker/structs"

	"expense-tracker/structs"
	"expense-tracker/utils"
	"fmt"

	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func GetUser(db *sql.DB, id int) (structs.User, error) {
	var user structs.User
	sql := "SELECT id, email FROM users where id = $1"

	err := db.QueryRow(sql, id).Scan(&user.ID, &user.Email)
	if err != nil {
		return user, nil
	}

	return user, err
}

func Register(db *sql.DB, user structs.User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	sql := "INSERT INTO USERS(email, password) VALUES ($1, $2) RETURNING id"

	var id int

	errs := db.QueryRow(sql, user.Email, hashedPassword).Scan(&id)
	if errs != nil {
		return 0, errs
	}

	return id, nil
}

func Login(db *sql.DB, credentials structs.Credentials) (string, error) {

	var user structs.User
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := db.QueryRow(query, credentials.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return "", fmt.Errorf("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", fmt.Errorf("Invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
