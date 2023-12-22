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
	uc := controllers.NewUserController(s)
	ac := controllers.NewActivityController(s)
	authMiddleware := middlewares.NewAuthMiddlewarer(db, s)

	r.Use(middlewares.Logging)

	r.HandleFunc("/auth/register", uc.RegisterUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", uc.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/mail", uc.MailCheckHandler).Methods(http.MethodGet)

	appRouter := r.PathPrefix("/problem").Subrouter()
	appRouter.Use(authMiddleware.Authorization)

	appRouter.HandleFunc("/list", pc.GetProblemListHandler).Methods(http.MethodGet)
	appRouter.HandleFunc("/{id:[0-9]+}", pc.GetProblemDetailHandler).Methods(http.MethodGet)
	appRouter.HandleFunc("/add", pc.PostProblemHandler).Methods(http.MethodPost)
	appRouter.HandleFunc("/like", ac.ChangeLikedHandler).Methods(http.MethodPost)
	// appRouter.HandleFunc("/solve", ac.).Methods(http.MethodPost)

	return r
}
