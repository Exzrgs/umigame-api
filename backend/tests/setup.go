package tests

import (
	"fmt"
)

func (t *Test) Setup() error {
	err := cleanupDB(t.env)
	if err != nil {
		fmt.Println("cleanup:", err)
		return err
	}

	err = setupTestData(t.env)
	if err != nil {
		fmt.Println("setup:", err)
		return err
	}

	return nil
}
