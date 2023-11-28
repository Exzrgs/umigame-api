package tasks_test

import (
	"testing"
	"database/sql"
	"os"
	"fmt"

	"umigame-api/utils"
	"umigame-api/tests"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var db *sql.DB
var test tests.Test

func TestMain(m *testing.M){
	db, err := utils.ConnectDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	test = tests.Test{DB: db}

	if err := test.Setup(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m.Run()

	if err := test.Teardown(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}