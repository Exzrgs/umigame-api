package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"umigame-api/models"
	"umigame-api/myerrors"
)

// @Summary 問題リストの取得
// @Description 問題リストの取得
// @Accept json
// @Produce json
// @Security UUID
// @Success 200
// @Failure 400
// @Router /problem/list [get]
func (c *ProblemController) GetProblemListHandler(w http.ResponseWriter, req *http.Request) {
	problems, err := c.service.GetProblemListService()
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(problems)
}

func (c *ProblemController) GetProblemDetailHandler(w http.ResponseWriter, req *http.Request) {
	var id int
	q := req.URL.Query()
	idStr, ok := q["id"]
	if ok && len(idStr) > 0 {
		var err error
		id, err = strconv.Atoi(idStr[0])
		if err != nil {
			err = myerrors.BadParameter.Wrap(err, "query parameter must be number")
			myerrors.ErrorHandler(w, req, err)
			return
		}
	}

	problem, err := c.service.GetProblemDetailService(id)
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(problem)
}

func (c *ProblemController) PostProblemHandler(w http.ResponseWriter, req *http.Request) {
	var problem models.Problem
	if err := json.NewDecoder(req.Body).Decode(&problem); err != nil {
		err = myerrors.ReqDecodeFailed.Wrap(err, "failed to decode json request body")
		myerrors.ErrorHandler(w, req, err)
	}

	newProblem, err := c.service.PostProblemService(problem)
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(newProblem)
}
