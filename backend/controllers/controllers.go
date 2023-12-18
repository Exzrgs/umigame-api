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

// type AuthController struct {
// 	service services.AuthServicer
// }

func NewProblemController(s services.ProblemServicer) *ProblemController {
	return &ProblemController{service: s}
}

// func NewAuthController(s services.AuthServicer) *AuthController {
// 	return &AuthController{service: s}
// }
