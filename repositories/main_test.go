package repositories_test

import (
	"database/sql"
	"os"
	"testing"
	"fmt"

	"umigame-api/utils"
	"umigame-api/tests"
)
/*
TODO:テストするときはproblem.goのページ数も変更しないといけない。大変マズイ
*/

var (
	db *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	db, err = utils.ConnectDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	test := tests.Test{DB: db}

	if err := test.Setup(); err != nil {
		os.Exit(1)
	}

	m.Run()

	if err := test.Teardown(); err != nil {
		os.Exit(1)
	}
}
