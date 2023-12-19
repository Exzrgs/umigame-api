package models

import (
	"time"
)

/*
Here is DB models
*/

type Problem struct {
	ID           int       `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Statement    string    `json:"statement" db:"statement"`
	Answer       string    `json:"answer" db:"answer"`
	Author       string    `json:"author" db:"author"`
	Reference    string    `json:"reference" db:"reference"`
	ReferenceURL string    `json:"reference_url" db:"reference_url"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type User struct {
	ID           int       `db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"password"`
	PasswordHash string    `db:"password_hash"`
	UUID         string    `db:"uuid"`
	IsValid      bool      `db:"is_valid"`
	CreatedAt    time.Time `db:"created_at"`
}

type Activity struct {
	UserID    int  `db:"user_id"`
	ProblemID int  `db:"problem_id"`
	IsSolved  bool `db:"is_solved"`
	IsLiked   bool `db:"is_liked"`
}

type Chat struct {
	ChatID    int
	ProblemID int
	UserID    int
	Question  string
	Answer    string
	CreatedAt time.Time
}
