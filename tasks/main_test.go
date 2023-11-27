package tasks_test

import (
	"testing"
	"database/sql"
	"os"
	"fmt"

	"umigame-api/utils"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var db *sql.DB

func TestMain(m *testing.M){
	db, err := utils.ConnectDB()
	if err != nil {
		os.Exit(1)
	}

	m.Run()
}