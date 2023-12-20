package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"umigame-api/models"
	"umigame-api/myerrors"
)

func (c *ProblemController) GetProblemListHandler(w http.ResponseWriter, req *http.Request) {
	var page int
	q := req.URL.Query()
	p, ok := q["page"]
	if ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = myerrors.BadParameter.Wrap(err, "query parameter must be number")
			myerrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		err := myerrors.BadParameter.Wrap(NoQueryParameter, "must have page query parameter")
		myerrors.ErrorHandler(w, req, err)
		return
	}

	problems, err := c.service.GetProblemListService(page)
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(problems)
}

func (c *ProblemController) GetProblemDetailHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = myerrors.BadParameter.Wrap(err, "query parameter must be number")
		myerrors.ErrorHandler(w, req, err)
		return
	}

	problem, err := c.service.GetProblemDetailService(id)
	if err != nil {
		myerrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(problem)
}

/*
	{
		"title": "test",
		"statement": "test",
		"answer": "test",
		"author": "test",
		"reference": "test",
		"reference_url":"test"
	}
*/
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
