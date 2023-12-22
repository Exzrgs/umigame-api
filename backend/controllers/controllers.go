package controllers

import (
	"umigame-api/services"
)

/*
mainやroutersから情報が欲しいときに使う
それ以外からは引数で受け取る
*/
type ProblemController struct {
	service services.ProblemServicer
}

type UserController struct {
	service services.UserServicer
}

type ActivityController struct {
	service services.ActivityServicer
}

func NewProblemController(s services.ProblemServicer) *ProblemController {
	return &ProblemController{service: s}
}

func NewUserController(s services.UserServicer) *UserController {
	return &UserController{service: s}
}

func NewActivityController(s services.ActivityServicer) *ActivityController {
	return &ActivityController{service: s}
}
