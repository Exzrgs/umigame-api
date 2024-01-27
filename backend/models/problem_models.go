package models

import (
	"time"
)

/*
GetProblemList : Response
*/
type ProblemBase struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Statement string    `json:"statement"`
	IsSolved  bool      `json:"is_solved"`
	IsLiked   bool      `json:"is_liked"`
	CreatedAt time.Time `json:"created_at"`
}

/*
GetProblemDetail : Response
PostProblem : Request
*/
type ProblemDetail struct {
	Title        string    `json:"title"`
	Statement    string    `json:"statement"`
	Answer       string    `json:"answer"`
	Author       string    `json:"author"`
	Reference    string    `json:"reference"`
	ReferenceURL string    `json:"reference_url"`
	CreatedAt    time.Time `json:"created_at"`

	IsSolved bool `json:"is_solved"`
	IsLiked  bool `json:"is_liked"`

	ChatList []ChatBase `json:"chat"`
}
