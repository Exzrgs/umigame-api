package routers

import (
	"database/sql"
	"net/http"

	"umigame-api/controllers"
	"umigame-api/middlewares"
	"umigame-api/services"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB, port string) *mux.Router {
	r := mux.NewRouter()

	s := services.NewServicer(db, port)
	pc := controllers.NewProblemController(s)
	ac := controllers.NewAuthController(s)

	r.HandleFunc("/problem/list", pc.GetProblemListHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/detail", pc.GetProblemDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/add", pc.PostProblemHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/register", ac.RegisterUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", ac.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/mail", ac.MailCheckHandler).Methods(http.MethodGet)
	// r.HandleFunc("/auth/cookie")

	r.Use(middlewares.Logging)

	return r
}
