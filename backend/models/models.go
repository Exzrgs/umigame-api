package models

import (
	"time"
)

type ProblemOutline struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	IsSolved  bool      `json:"is_solved"`
	CreatedAt time.Time `json:"created_at"`
}

type ProblemDetail struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	ProblemStatement string    `json:"problem_statement"`
	ProblemAnswer    string    `json:"problem_answer"`
	UserChat         []string  `json:"user_chat"`
	GPTChat          []string  `json:"gpt_chat"`
	IsSolved         bool      `json:"is_solved"`
	CreatedAt        time.Time `json:"created_at"`
}

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	SolvedList []int     `json:"solved_list"`
	CreatedAt  time.Time `json:"created_at"`
}

type Auth struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Hash         string
	Uuid         string
	ActivateFlag bool
}

type Uuid struct {
	Uuid string `json:"uuid"`
}
