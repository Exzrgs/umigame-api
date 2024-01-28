package tests

import (
	"fmt"
)

func (t *Test) Setup() error {
	err := cleanupDB(t.env)
	if err != nil {
		fmt.Println("cleanup:", err.Error())
		return err
	}

	err = setupTestData(t.env)
	if err != nil {
		fmt.Println("setup:", err.Error())
		return err
	}

	return nil
}
