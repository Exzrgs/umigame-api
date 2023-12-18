package repositories

import (
	"umigame-api/models"
	"umigame-api/myerrors"

	"github.com/jmoiron/sqlx"
)

func SelectChatDetail(db *sqlx.DB, userID int, problemID int) ([]models.Chat, error) {
	sqlStr := `
	SELECT question, answer, created_at
	FROM chats
	WHERE user_id = ? AND problem_id = ?;
	`

	rows, err := db.Queryx(sqlStr, userID, problemID)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}
	defer rows.Close()

	chatList := make([]models.Chat, 0)

	for rows.Next() {
		var chat models.Chat
		if err := rows.StructScan(&chat); err != nil {
			err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
			return nil, err
		}
		chatList = append(chatList, chat)
	}

	return chatList, nil
}

func SelectChatList(db *sqlx.DB, userID int) ([]models.Chat, error) {
	sqlStr := `
	SELECT problem_id, question, created_at
	FROM chats
	WHERE user_id = ?;
	`

	rows, err := db.Queryx(sqlStr, userID)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}
	defer rows.Close()

	chatList := make([]models.Chat, 0)

	for rows.Next() {
		var chat models.Chat
		if err := rows.StructScan(&chat); err != nil {
			err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
			return nil, err
		}
		chatList = append(chatList, chat)
	}

	return chatList, nil
}
