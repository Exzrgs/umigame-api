package routers

import (
	"database/sql"
	"net/http"

	"umigame-api/controllers"
	"umigame-api/middlewares"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	pc := controllers.NewProblemController(db)
	ac := controllers.NewAuthController(db)

	r.HandleFunc("/problem/list", pc.GetProblemListHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/detail", pc.GetProblemDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/problem/add", pc.PostProblemHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/register", ac.RegisterUserHandler).Methods(http.MethodPost)
	// r.HandleFunc("/auth/login")
	r.HandleFunc("/auth/mail", ac.MailCheckHandler).Methods(http.MethodGet)
	// r.HandleFunc("/auth/cookie")

	r.Use(middlewares.Logging)

	return r
}
