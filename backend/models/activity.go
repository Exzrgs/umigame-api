package models

type ActivityLiked struct {
	UserID    int  `json:"user_id"`
	ProblemID int  `json:"problem_id"`
	IsLiked   bool `json:"is_liked"`
}
