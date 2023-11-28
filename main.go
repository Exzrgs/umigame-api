package main

import (
	"database/sql"
	"log"
	"net/http"

	"umigame-api/routers"
	"umigame-api/utils"
	"umigame-api/tasks"

	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

var (
	db         *sql.DB
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbDatabase string
	dbConfig   string
)

func main() {
	db, err := utils.ConnectDB()
	if err != nil {
		log.Println(err)
		return
	}

	r := routers.NewRouter(db)

	go tasks.ExeTasks(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
