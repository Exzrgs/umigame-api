package models

import "time"

type ChatBase struct {
	Question  string
	Answer    string
	CreatedAt time.Time
}

type ChatRoom struct {
	ProblemID    int
	ProblemTitle string
	Question     string
	CreatedAt    time.Time
}
