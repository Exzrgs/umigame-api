package utils

import (
	"fmt"

	"umigame-api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(env models.Env) (*sqlx.DB, error) {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.DBDatabase)

	db, err := sqlx.Open("mysql", dbConfig)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
