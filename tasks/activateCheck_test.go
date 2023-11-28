package tasks_test

import (
	"testing"
	"database/sql"

	"umigame-api/tasks"
)

func TestActivateCheck(t *testing.T){
	bef, err := countData(test.DB)
	if err != nil {
		t.Fatal(err)
	}

	tasks.ActivateCheck(test.DB)

	aft, err := countData(test.DB)
	if err != nil {
		t.Fatal(err)
	}

	if bef == 0 {
		t.Errorf("after 15 minutes data should exist before exe, but 0\n")
	}
	if aft > 0 {
		t.Errorf("after 15 minutes data exist after exe, %d count\n", aft)
	}
}

/*
DBから特定のデータを取り出し、その数を返す
*/
func countData(db *sql.DB)(int, error){
	const sqlStr = `
	select count(*) from users
	where activate_flag = false and timestampdiff(minute, created_at, now()) >= 15;
	`
	rows, err := db.Query(sqlStr)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}

	return count, nil
}