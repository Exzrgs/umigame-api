package repositories_test

import (
	"fmt"
	"os"
	"testing"

	"umigame-api/models"
	"umigame-api/tests"
	"umigame-api/utils"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

/*
TODO:テストするときはproblem.goのページ数も変更しないといけない。大変マズイ
*/

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	db *sqlx.DB
)

func TestMain(m *testing.M) {
	var env models.Env
	envconfig.Process("test", &env)

	var err error
	db, err = utils.ConnectDB(env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	test := tests.NewTest(db, env)

	if err := test.Setup(); err != nil {
		os.Exit(1)
	}

	m.Run()

	if err := test.Teardown(); err != nil {
		os.Exit(1)
	}
}
