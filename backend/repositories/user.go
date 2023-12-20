package repositories

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"umigame-api/models"
	"umigame-api/myerrors"
)

func RegisterUser(db *sqlx.DB, user *models.User) error {
	const sqlStr = `
	INSERT INTO users (name, email, password_hash, uuid, is_valid, created_at) VALUES
	(?, ?, ?, ?, false, now());
	`

	if _, err := db.Exec(sqlStr, user.Name, user.Email, user.PasswordHash, user.UUID); err != nil {
		/*
			TODO:emailが被ってるのか、普通にエラーなのかで分岐する
		*/
		err = myerrors.InsertDataFailed.Wrap(err, "internal server error")
		return err
	}

	return nil
}

func UpdateActivate(db *sqlx.DB, uuid string) error {
	const sqlStr = `
	UPDATE users
	SET is_valid = TRUE
	WHERE uuid = ?;
	`

	result, err := db.Exec(sqlStr, uuid)
	if err != nil {
		err = myerrors.UpdateDataFailed.Wrap(err, "internal server error")
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		err = myerrors.Unknown.Wrap(err, "internal server error")
		return err
	}
	if num == 0 {
		err = myerrors.NoData.Wrap(ErrNotFound, "not found")
		return err
	}

	return nil
}

func GetAuthInfo(db *sqlx.DB, email string) (models.User, error) {
	const sqlStr = `
	SELECT password_hash, uuid, is_valid
	FROM users
	WHERE email = ?;
	`

	var user models.User
	if err := db.QueryRowx(sqlStr, email).StructScan(&user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = myerrors.NoData.Wrap(err, "email is not registered")
			return models.User{}, err
		} else {
			err = myerrors.GetDataFailed.Wrap(err, "internal server error")
			return models.User{}, err
		}
	}

	return user, nil
}
