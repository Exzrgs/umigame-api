package tests

import (
	"fmt"
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