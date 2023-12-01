package controllers

import (
	"encoding/json"
	"net/http"

	"umigame-api/models"
	"umigame-api/myerrors"
)

func (c *AuthController) RegisterUserHandler(w http.ResponseWriter, req *http.Request) {
	var auth models.Auth
	if err := json.NewDecoder(req.Body).Decode(&auth); err != nil {
		err = myerrors.ReqDecodeFailed.Wrap(err, "failed to decode json request body")
		myerrors.ErrorHandler(w, req, err)
		return
	}

	if err := c.service.RegisterUserService(auth); err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *AuthController) MailCheckHandler(w http.ResponseWriter, req *http.Request) {
	uuid := req.URL.Query().Get("uuid")

	if err := c.service.MailCheckService(uuid); err != nil {
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

	uuid, err := c.service.LoginService(auth.Email, auth.Password)
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(models.Uuid{Uuid: uuid})
}
