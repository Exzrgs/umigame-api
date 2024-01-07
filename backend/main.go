package main

import (
	"flag"
	"log"
	"net/http"

	"umigame-api/routers"
	"umigame-api/tasks"
	"umigame-api/utils"
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
