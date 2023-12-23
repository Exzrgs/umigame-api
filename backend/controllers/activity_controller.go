package controllers

import (
	"encoding/json"
	"net/http"
	"umigame-api/models"
	"umigame-api/myerrors"
)

/*
連打されても大丈夫なように、フロントからもらった状態を反転したやつに変更する
*/
func (c *ActivityController) ChangeLikedHandler(w http.ResponseWriter, req *http.Request) {
	var activity models.Activity
	if err := json.NewDecoder(req.Body).Decode(&activity); err != nil {
		err = myerrors.ReqDecodeFailed.Wrap(err, "failed to decode json request body")
		myerrors.ErrorHandler(w, req, err)
	}

	newActivity, err := c.service.ChangeLikedService(activity)
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(newActivity)
}
