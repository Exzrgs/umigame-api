package main

import (
	"flag"
	"log"
	"net/http"

	"umigame-api/models"
	"umigame-api/routers"
	"umigame-api/tasks"
	"umigame-api/utils"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	port := flag.String("p", ":8080", "HTTP network port")
	flag.Parse()

	var env models.Env
	envconfig.Process("", &env)

	db, err := utils.ConnectDB(env)
	if err != nil {
		log.Println(err)
		return
	}

	r := routers.NewRouter(db, *port)

	go tasks.ExeTasks(db)

	log.Printf("server start at port %s", *port)
	log.Fatal(http.ListenAndServe(*port, r))
}
