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
	c := controllers.NewController(s, db)

	r.HandleFunc("/problem/list", c.GetProblemListHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/detail", c.GetProblemDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/add", c.PostProblemHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/register", c.RegisterUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", c.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/mail", c.MailCheckHandler).Methods(http.MethodGet)
	// r.HandleFunc("/auth/cookie")

	r.Use(middlewares.Logging)

	return r
}
