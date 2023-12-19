package services

import (
	"umigame-api/models"

	"github.com/jmoiron/sqlx"
)

type ProblemServicer interface {
	GetProblemListService(page int) ([]models.ProblemBase, error)
	GetProblemDetailService(problemID int) (models.ProblemDetail, error)
	PostProblemService(problem models.Problem) (models.Problem, error)
}

// type AuthServicer interface {
// 	RegisterUserService(auth models.Auth) error
// 	MailCheckService(uuid string) error
// 	LoginService(email string, password string) (string, error)
// }

/*
mainやroutersから情報が欲しいときに使う
それ以外からは引数で受け取る
*/
type Service struct {
	db     *sqlx.DB
	userID int
	port   string
}

func NewServicer(db *sqlx.DB, port string) *Service {
	return &Service{
		db:   db,
		port: port,
	}
}

func (s *Service) SetUserID(id int) {
	s.userID = id
}
