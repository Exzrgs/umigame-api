package tests

import (
	"fmt"
	"os/exec"

	"umigame-api/models"

	"github.com/jmoiron/sqlx"
)

type Test struct {
	db  *sqlx.DB
	env models.Env
}

func NewTest(db *sqlx.DB, env models.Env) *Test {
	return &Test{db: db, env: env}
}

func setupTestData(env models.Env) error {
	cmd := exec.Command("mysql", "-h", env.DBHost, "-u", env.DBUser, env.DBDatabase, fmt.Sprintf("--password=%s", env.DBPassword), "-e", "source ./testdata/setupDB.sql")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func cleanupDB(env models.Env) error {
	cmd := exec.Command("mysql", "-h", env.DBHost, "-u", env.DBUser, env.DBDatabase, fmt.Sprintf("--password=%s", env.DBPassword), "-e", "source ./testdata/cleanupDB.sql")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
