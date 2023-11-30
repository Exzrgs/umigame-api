package tests

func (t *Test) Teardown() error {
	err := cleanupDB()
	if err != nil {
		return err
	}

	t.DB.Close()

	return nil
}