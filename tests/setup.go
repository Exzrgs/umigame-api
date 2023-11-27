package tests

import (
	"os/exec"
	"fmt"
	// "os"

	//"umigame-api/utils"
	_ "github.com/go-sql-driver/mysql"
)

func (t *Test) Setup() error {
	err := cleanupDB()
	if err != nil {
		fmt.Println("cleanup:", err)
		return err
	}

	err = setupTestData()
	if err != nil {
		fmt.Println("setup:", err)
		return err
	}

	return nil
}

func setupTestData() error {
	cmd := exec.Command("mysql", "-h", dbHost, "-u", dbUser, dbDatabase, fmt.Sprintf("--password=%s", dbPassword), "-e", "source ./testdata/setupDB.sql")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}