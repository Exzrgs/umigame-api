package controllers

import (
	"encoding/json"
	"net/http"

	"umigame-api/models"
	"umigame-api/myerrors"
)

func (c *UserController) RegisterUserHandler(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		err = myerrors.ReqDecodeFailed.Wrap(err, "failed to decode json request body")
		myerrors.ErrorHandler(w, req, err)
		return
	}

	if err := c.service.RegisterUserService(user); err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *UserController) MailCheckHandler(w http.ResponseWriter, req *http.Request) {
	uuid := req.URL.Query().Get("uuid")

	if err := c.service.MailCheckService(uuid); err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *UserController) LoginHandler(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		err = myerrors.ReqDecodeFailed.Wrap(err, "failed to decode json request body")
		myerrors.ErrorHandler(w, req, err)
		return
	}

	uuid, err := c.service.LoginService(user.Email, user.Password)
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(models.User{UUID: uuid})
}
