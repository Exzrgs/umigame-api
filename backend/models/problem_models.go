package models

import (
	"time"
)

/*
GetProblemList : Response
*/
type ProblemBase struct {
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
	ChatList     []ChatBase
	CreatedAt    time.Time
}

type ChatBase struct {
	Question  string
	Answer    string
	CreatedAt time.Time
}
