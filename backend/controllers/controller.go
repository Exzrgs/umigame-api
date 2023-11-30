package controllers

import (
	"database/sql"

	"umigame-api/services"
)

/*
mainやroutersから情報が欲しいときに使う
それ以外からは引数で受け取る
*/
type Controller struct {
	service *services.Servicer
	db      *sql.DB
}

func NewController(s *services.Servicer, db *sql.DB)*Controller{
	return &Controller{
		service: s,
		db: db,
	}
}