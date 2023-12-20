package repositories

import (
	"umigame-api/models"
	"umigame-api/myerrors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
)

const (
	/*
		テストするときは、ここも変更する必要がある.
		テストケースを増やしてこのままテストできるようにしたい
	*/
	ProblemNumPerPage = 20
)

func SelectProblemList(db *sqlx.DB, page int) ([]models.Problem, error) {
	const sqlStr = `
	SELECT id, title, author, statement, created_at
	FROM problems
	LIMIT ? OFFSET ?;
	`

	rows, err := db.Queryx(sqlStr, ProblemNumPerPage, (page-1)*ProblemNumPerPage)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}
	defer rows.Close()

	problemList := make([]models.Problem, 0)

	for rows.Next() {
		var problem models.Problem
		if err := rows.StructScan(&problem); err != nil {
			err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
			return nil, err
		}
		problemList = append(problemList, problem)
	}

	return problemList, nil
}

func SelectProblem(db *sqlx.DB, ID int) (models.Problem, error) {
	sqlStr := `
	SELECT title, statement, answer, author, reference, reference_url, created_at
	FROM problems
	WHERE id = ?;
	`

	var problem models.Problem
	row := db.QueryRowx(sqlStr, ID)
	if err := row.Err(); err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.Problem{}, err
	}

	if err := row.StructScan(&problem); err != nil {
		err = myerrors.NoData.Wrap(err, "problem is not found")
		return models.Problem{}, err
	}

	return problem, nil
}

func InsertProblem(db *sqlx.DB, problem models.Problem) (models.Problem, error) {
	const sqlStr = `
	INSERT INTO problems (title, statement, answer, author, reference, reference_url, created_at) VALUES
	(?, ?, ?, ?, ?, ?, now());
	`

	var newProblem models.Problem
	if err := copier.Copy(&newProblem, &problem); err != nil {
		err = myerrors.TypeCastFailed.Wrap(err, "internal server error")
		return models.Problem{}, err
	}

	result, err := db.Exec(sqlStr, problem.Title, problem.Statement, problem.Answer, problem.Author, problem.Reference, problem.ReferenceURL)
	if err != nil {
		err = myerrors.InsertDataFailed.Wrap(err, "failed to insert data")
		return models.Problem{}, err
	}

	id, _ := result.LastInsertId()
	newProblem.ID = int(id)

	return newProblem, nil
}
