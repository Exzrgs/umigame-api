package models

import (
	"time"

	"github.com/google/uuid"
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
	Password     string
	Hash         string
	Uuid         uuid.UUID
	ActivateFlag bool
	SolvedList   []int
	CreatedAt    time.Time
}

type Chat struct {
	ChatID    int
	ProblemID int
	UserID    int
	Question  string
	Answer    string
	CreatedAt time.Time
}
