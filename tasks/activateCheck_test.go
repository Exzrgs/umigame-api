package tasks_test

func TestactivateCheck(t *testing){
	t.Run(func(t *testing){
		bef, err := countData()
		if err != nil {
			t.Fatal(err)
		}

		activateCheck()

		aft, err := countData()
		if err != nil {
			t.Fatal(err)
		}

		if bef == 0 {
			t.Errorf("after 15 minutes data should exist before exe, but 0\n")
		}
		if aft > 0 {
			t.Errorf("after 15 minutes data exist after exe, %d count\n", aft)
		}
	})
}

/*
DBから特定のデータを取り出し、その数を返す
*/
func countData()(int, error){
	const sqlStr = `
	select * from users
	where activate_flag = false and timestampdiff(minute, created_at, now()) >= 15;
	`
	rows, err := db.Query(sqlStr)
	if err != nil {
		return 0, err
	}

	return len(rows), nil
}