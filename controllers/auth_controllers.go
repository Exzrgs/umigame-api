package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"umigame-api/models"
	"umigame-api/myerrors"
	"umigame-api/services"
)

type AuthController struct {
	db *sql.DB
}

func NewAuthController(db *sql.DB) *AuthController {
	return &AuthController{db: db}
}

func (c *AuthController) RegisterUserHandler(w http.ResponseWriter, req *http.Request) {
	var auth models.Auth
	if err := json.NewDecoder(req.Body).Decode(&auth); err != nil {
		err = myerrors.ReqDecodeFailed.Wrap(err, "failed to decode json request body")
		myerrors.ErrorHandler(w, req, err)
		return
	}

	if err := services.RegisterUserService(c.db, &auth); err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *AuthController) MailCheckHandler(w http.ResponseWriter, req *http.Request) {
	uuid := req.URL.Query().Get("uuid")

	if err := services.MailCheckService(c.db, uuid); err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *AuthController) LoginHandler(w http.ResponseWriter, req *http.Request) {
	var auth models.Auth
	if err := json.NewDecoder(req.Body).Decode(&auth); err != nil {
		err = myerrors.ReqDecodeFailed.Wrap(err, "failed to decode json request body")
		myerrors.ErrorHandler(w, req, err)
		return
	}
	
	uuid, err := services.LoginService(c.db, auth.Email, auth.Password)
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(models.Uuid{Uuid: uuid})
}
