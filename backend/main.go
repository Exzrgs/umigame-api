package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"umigame-api/routers"
	"umigame-api/tasks"
	"umigame-api/utils"

	"github.com/joho/godotenv"
)

func init() {
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
	port := flag.String("p", ":8080", "HTTP network port")
	flag.Parse()

	db, err := utils.ConnectDB()
	if err != nil {
		log.Println(err)
		return
	}

	r := routers.NewRouter(db, *port)

	go tasks.ExeTasks(db)

	log.Printf("server start at port %s", *port)
	log.Fatal(http.ListenAndServe(*port, r))
}
