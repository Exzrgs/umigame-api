package tasks

import (
	"database/sql"
	"log"

	"umigame-api/myerrors"
)

/*
1分起きにactivate_flagをチェック
TODO: ユニットテスト
*/
func activateCheck(db *sql.DB) {
	const sqlStr = `
	delete from users
	where activate_flag = false and timestampdiff(minute, created_at, now()) >= 15;
	`

	_, err := db.Exec(sqlStr)
	if err != nil {
		myerrors.DeleteDataFailed.Wrap(err, "internal server error")
		log.Println(err)
	}
}
