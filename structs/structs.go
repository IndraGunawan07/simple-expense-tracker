package structs

import "time"

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Category struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type Expense struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	Types       int       `json:"types"`
	Dates       time.Time `json:"dates"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type InsertExpense struct {
	CategoryID  int    `json:"category_id" binding:"required"`
	Types       int    `json:"types" binding:"required"`
	Dates       string `json:"dates" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
	Description string `json:"description"`
}
