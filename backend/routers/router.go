package routers

import (
	"net/http"

	"umigame-api/controllers"
	"umigame-api/middlewares"
	"umigame-api/services"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB, port string) *mux.Router {
	r := mux.NewRouter()
	s := services.NewServicer(db, port)
	pc := controllers.NewProblemController(s)
	// ac := controllers.NewAuthController(s)
	a := middlewares.NewAuthMiddlewarer(db, s)

	r.Use(middlewares.Logging)

	// r.HandleFunc("/auth/register", ac.RegisterUserHandler).Methods(http.MethodPost)
	// r.HandleFunc("/auth/login", ac.LoginHandler).Methods(http.MethodPost)
	// r.HandleFunc("/auth/mail", ac.MailCheckHandler).Methods(http.MethodGet)

	r.Use(a.Authorization)

	r.HandleFunc("/problem/list", pc.GetProblemListHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/detail", pc.GetProblemDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/add", pc.PostProblemHandler).Methods(http.MethodPost)

	return r
}
