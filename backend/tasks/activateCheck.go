package tasks

import (
	"log"

	"umigame-api/myerrors"

	"github.com/jmoiron/sqlx"
)

/*
1分起きにactivate_flagをチェック
TODO: ユニットテスト この関数を実行した後に、15分以上のやつがなくなっていること
*/
func ActivateCheck(db *sqlx.DB) {
	const sqlStr = `
	delete from users
	where is_valid = false and timestampdiff(minute, created_at, now()) >= 15;
	`

	_, err := db.Exec(sqlStr)
	if err != nil {
		myerrors.DeleteDataFailed.Wrap(err, "internal server error")
		log.Println(err)
	}
}
