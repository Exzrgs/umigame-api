package models

import (
	"time"
)

type Problem struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	ProblemStatement string    `json:"problem_statement ,omitempty"`
	ProblemAnswer    string    `json:"problem_answer ,omitempty"`
	UserChat         []string  `json:"user_chat ,omitempty"`
	GPTChat          []string  `json:"gpt_chat ,omitempty"`
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
	Email        string `json:"email ,omitempty"`
	Password     string `json:"password ,omitempty"`
	Hash         string `json:"-"`
	Uuid         string `json:"-"`
	ActivateFlag bool   `json:"-"`
}
