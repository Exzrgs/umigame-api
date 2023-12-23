package repositories

import (
	"umigame-api/models"
	"umigame-api/myerrors"

	"github.com/jmoiron/sqlx"
)

func SelectProblemChat(db *sqlx.DB, userID int, problemID int) ([]models.Chat, error) {
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

func SelectChatroomList(db *sqlx.DB, userID int) ([]models.Chatroom, error) {
	sqlStr := `
    SELECT problem_id, problems.title, c.question, c.created_at FROM (
        SELECT problem_id, question, created_at FROM chats WHERE 
            id IN (
                SELECT MAX(id) FROM chats WHERE user_id = ? GROUP BY problem_id)
            )
    AS c INNER JOIN problems ON problems.id = c.problem_id;
	`

	rows, err := db.Queryx(sqlStr, userID)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}
	defer rows.Close()

	chatroomList := make([]models.Chatroom, 0)

	for rows.Next() {
		var chatroom models.Chatroom
		if err := rows.StructScan(&chatroom); err != nil {
			err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
			return nil, err
		}
		chatroomList = append(chatroomList, chatroom)
	}

	return chatroomList, nil
}
