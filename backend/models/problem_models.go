package models

import (
	"time"
)

/*
GetProblemList : Response
*/
type ProblemBase struct {
	ID        int
	Title     string
	Author    string
	Statement string
	IsSolved  bool
	IsLiked   bool
	CreatedAt time.Time
}

/*
GetProblemDetail : Response
PostProblem : Request
*/
type ProblemDetail struct {
	Title        string
	Statement    string
	Answer       string
	Author       string
	Reference    string
	ReferenceURL string
	CreatedAt    time.Time

	IsSolved bool
	IsLiked  bool

	ChatList []ChatBase
}
