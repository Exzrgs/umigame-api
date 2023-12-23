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
	cc := controllers.NewChatController(s)
	authMiddleware := middlewares.NewAuthMiddlewarer(db, s)

	r.Use(middlewares.Logging)

	r.HandleFunc("/auth/register", uc.RegisterUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", uc.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/mail", uc.MailCheckHandler).Methods(http.MethodGet)

	problemRouter := r.PathPrefix("/problem").Subrouter()
	problemRouter.Use(authMiddleware.Authorization)

	problemRouter.HandleFunc("/list", pc.GetProblemListHandler).Methods(http.MethodGet)
	problemRouter.HandleFunc("/{id:[0-9]+}", pc.GetProblemDetailHandler).Methods(http.MethodGet)
	problemRouter.HandleFunc("/add", pc.PostProblemHandler).Methods(http.MethodPost)
	problemRouter.HandleFunc("/like", ac.ChangeLikedHandler).Methods(http.MethodPost)
	// appRouter.HandleFunc("/solve", ac.).Methods(http.MethodPost)

	chatRouter := r.PathPrefix("/chat").Subrouter()
	chatRouter.Use(authMiddleware.Authorization)

	chatRouter.HandleFunc("/room/list", cc.GetChatroomListHandler).Methods(http.MethodGet)
	// chatRouter.HandleFunc("/room").Methods(http.MethodPost)
	// chatRouter.HandleFunc("/room").Methods(http.MethodDelete)

	return r
}
