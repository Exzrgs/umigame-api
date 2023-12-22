package repositories

import (
	"umigame-api/models"
	"umigame-api/myerrors"

	"github.com/jmoiron/sqlx"
)

func SelectActivityList(db *sqlx.DB, userID int, page int, problemIDs []int) ([]models.Activity, error) {
	sqlStr := `
	SELECT problem_id, is_solved, is_liked
	FROM activities
	WHERE user_id = ? AND problem_id IN (?);
	`

	sqlStr, params, err := sqlx.In(sqlStr, userID, problemIDs)
	if err != nil {
		myerrors.SQLPrepareFailed.Wrap(err, "internal server error")
		return nil, err
	}

	var activities []models.Activity
	rows, err := db.Queryx(sqlStr, params...)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var activity models.Activity
		if err := rows.StructScan(&activity); err != nil {
			err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
			return nil, err
		}
		activities = append(activities, activity)
	}

	return activities, nil
}

func SelectActivity(db *sqlx.DB, userID int, problemID int) (models.Activity, error) {
	sqlStr := `
	SELECT is_solved, is_liked
	FROM activities
	WHERE user_id = ? AND problem_id = ?;
	`

	rows, err := db.Queryx(sqlStr, userID, problemID)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.Activity{}, err
	}
	defer rows.Close()

	var activity models.Activity
	for rows.Next() {
		if err := rows.StructScan(&activity); err != nil {
			err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
			return models.Activity{}, err
		}
	}

	return activity, nil
}
