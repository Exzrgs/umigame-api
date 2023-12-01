package controllers

import (
	"umigame-api/services"
)

/*
mainやroutersから情報が欲しいときに使う
それ以外からは引数で受け取る
*/
type Controller struct {
	service *services.Servicer
}

func NewController(s *services.Servicer)*Controller{
	return &Controller{
		service: s,
	}
}