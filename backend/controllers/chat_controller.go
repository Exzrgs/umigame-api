package controllers

import (
	"encoding/json"
	"net/http"
	"umigame-api/models"
	"umigame-api/myerrors"
)

func (c *ChatController) GetChatroomListHandler(w http.ResponseWriter, req *http.Request) {
	chatrooms, err := c.service.GetChatroomListService()
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(chatrooms)
}

func (c *ChatController) PostQuestionHandler(w http.ResponseWriter, req *http.Request) {
	var questionChat models.Chat
	if err := json.NewDecoder(req.Body).Decode(&questionChat); err != nil {
		err = myerrors.ReqDecodeFailed.Wrap(err, "failed to decode json request body")
		myerrors.ErrorHandler(w, req, err)
		return
	}

	answerChat, err := c.service.PostQuestionService(questionChat)
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(answerChat)
}
