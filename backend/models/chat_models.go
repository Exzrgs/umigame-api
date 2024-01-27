package models

import "time"

type ChatBase struct {
	Question  string    `json:"question" db:"question"`
	Answer    string    `json:"answer" db:"answer"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Chatroom struct {
	ProblemID    int       `json:"problem_id" db:"problem_id"`
	ProblemTitle string    `json:"problem_title" db:"title"`
	LastQuestion string    `json:"last_question" db:"question"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
