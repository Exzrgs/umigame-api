package services

import (
	"database/sql"
)

/*
mainやroutersから情報が欲しいときに使う
それ以外からは引数で受け取る
*/
type Servicer struct {
	db   *sql.DB
	port string
}

func NewServicer(db *sql.DB, port string)*Servicer{
	return &Servicer{
		db: db,
		port: port,
	}
}