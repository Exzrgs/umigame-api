package models

import (
	"time"
)

/*
Here are DB models
*/

type Problem struct {
	ID           int        `json:"id,omitempty" db:"id"`
	Title        string     `json:"title,omitempty" db:"title"`
	Statement    string     `json:"statement,omitempty" db:"statement"`
	Answer       string     `json:"answer,omitempty" db:"answer"`
	Author       string     `json:"author,omitempty" db:"author"`
	Reference    string     `json:"reference,omitempty" db:"reference"`
	ReferenceURL string     `json:"reference_url,omitempty" db:"reference_url"`
	CreatedAt    *time.Time `json:"created_at,omitempty" db:"created_at"` // ポインタにするとnilになるからomitemptyが機能する
}

type User struct {
	ID           int        `json:"id,omitempty" db:"id"`
	Name         string     `json:"name,omitempty" db:"name"`
	Email        string     `json:"email,omitempty" db:"email"`
	Password     string     `json:"password,omitempty"`
	PasswordHash string     `json:"-" db:"password_hash"`
	UUID         string     `json:"uuid,omitempty" db:"uuid"`
	IsValid      bool       `json:"is_valid,omitempty" db:"is_valid"`
	CreatedAt    *time.Time `json:"created_at,omitempty" db:"created_at"`
}

type Activity struct {
	UserID    int  `json:"user_id,omitempty" db:"user_id"`
	ProblemID int  `json:"problem_id,omitempty" db:"problem_id"`
	IsSolved  bool `json:"is_solved,omitempty" db:"is_solved"`
	IsLiked   bool `json:"is_liked,omitempty" db:"is_liked"`
}

type Chat struct {
	ID        int        `json:"id,omitempty" db:"id"`
	ProblemID int        `json:"problem_id,omitempty" db:"problem_id"`
	UserID    int        `json:"user_id,omitempty" db:"user_id"`
	Question  string     `json:"question,omitempty" db:"question"`
	Answer    *int       `json:"answer,omitempty" db:"answer"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
}

type Env struct {
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     string `envconfig:"DB_PORT"`
	DBDatabase string `envconfig:"DB_DATABASE"`
}
