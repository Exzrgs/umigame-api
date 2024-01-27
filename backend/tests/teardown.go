package tests

func (t *Test) Teardown() error {
	err := cleanupDB(t.env)
	if err != nil {
		return err
	}

	t.db.Close()

	return nil
}
