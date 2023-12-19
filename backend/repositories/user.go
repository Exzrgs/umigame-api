package repositories

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"umigame-api/models"
	"umigame-api/myerrors"
)

func RegisterUser(db *sql.DB, user *models.User) error {
	const sqlStr = `
	INSERT INTO users (email, password_hash, uuid, activate_flag, created_at) values
	(?, ?, ?, false, now());
	`

	if _, err := db.Exec(sqlStr, user.Email, user.PasswordHash, user.UUID); err != nil {
		err = myerrors.InsertDataFailed.Wrap(err, "email is already used")
		return err
	}

	return nil
}

func UpdateActivate(db *sql.DB, uuid string) error {
	const sqlStr = `
	update users
	set activate_flag = true
	where uuid = ?;
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
		err = myerrors.NoData.Wrap(NotFound, "not found")
		return err
	}

	return nil
}

func GetAuthInfo(db *sqlx.DB, email string) (models.User, error) {
	const sqlStr = `
	select password_hash, uuid, activate_flag from users
	where email = ?;
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
