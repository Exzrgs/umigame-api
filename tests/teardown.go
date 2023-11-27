package tests

import (
	"os/exec"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func (t *Test) Teardown() error {
	err := cleanupDB()
	if err != nil {
		return err
	}

	t.DB.Close()

	return nil
}

func cleanupDB() error {
	cmd := exec.Command("mysql", "-h", dbHost, "-u", dbUser, dbDatabase, fmt.Sprintf("--password=%s", dbPassword), "-e", "source ./testdata/cleanupDB.sql")

	var err error
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}