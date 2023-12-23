package controllers

import (
	"encoding/json"
	"net/http"
	"umigame-api/myerrors"
)

func (c *ChatController) GetChatroomListHandler(w http.ResponseWriter, req *http.Request) {
	chatrooms, err := c.service.GetChatroomListService()
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
	}

	json.NewEncoder(w).Encode(chatrooms)
}
