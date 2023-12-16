package models

import (
	"time"
)

/*
Here is DB models
*/

type Problem struct {
	ID           int
	Title        string
	Statement    string
	Answer       string
	Author       string
	Reference    string
	ReferenceURL string
	CreatedAt    time.Time
}

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	UUID         string
	IsValid      bool
	CreatedAt    time.Time
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
