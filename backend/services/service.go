package services

import (
	"database/sql"

	"umigame-api/models"
)

type ProblemServicer interface {
	GetProblemListService(page int) ([]models.ProblemOutline, error)
	GetProblemDetailService(id int) (models.ProblemDetail, error)
	PostProblemService(problem models.ProblemDetail) (models.ProblemDetail, error)
}

type AuthServicer interface {
	RegisterUserService(auth models.Auth) error
	MailCheckService(uuid string) error
	LoginService(email string, password string) (string, error)
}

/*
mainやroutersから情報が欲しいときに使う
それ以外からは引数で受け取る
*/
type Servicer struct {
	db   *sql.DB
	port string
}

func NewServicer(db *sql.DB, port string) *Servicer {
	return &Servicer{
		db:   db,
		port: port,
	}
}
