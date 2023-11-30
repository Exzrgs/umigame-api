package tests

import (
	"database/sql"
	"os"
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

var (
	db         *sql.DB
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbDatabase string
	dbConfig   string
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbConfig = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbDatabase)
}

type Test struct {
	DB *sql.DB
}

func setupTestData() error {
	cmd := exec.Command("mysql", "-h", dbHost, "-u", dbUser, dbDatabase, fmt.Sprintf("--password=%s", dbPassword), "-e", "source ./testdata/setupDB.sql")

	err := cmd.Run()
	if err != nil {
		return err
	}

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