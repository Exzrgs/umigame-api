package models

import "time"

type ChatBase struct {
	Question  string
	Answer    string
	CreatedAt time.Time
}

type Chatroom struct {
	ProblemID    int       `json:"problem_id" db:"problem_id"`
	ProblemTitle string    `json:"problem_title" db:"title"`
	LastQuestion string    `json:"last_question" db:"question"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
