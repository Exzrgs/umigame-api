package middlewares

import (
	"github.com/jmoiron/sqlx"
)

func authorize(db *sqlx.DB, uuid string) (int, error) {
	sqlStr := `
	SELECT id
	FROM users
	WHERE uuid = ?;
	`

	var userID int
	if err := db.QueryRowx(sqlStr, uuid).Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}
