package repositories

import (
	"database/sql"

	"umigame-api/models"
	"umigame-api/myerrors"

	_ "github.com/go-sql-driver/mysql"
)

const (
	/*
		テストするときは、ここも変更する必要がある.
		テストケースを増やしてこのままテストできるようにしたい
	*/
	ProblemNumPerPage = 20
)

func SelectProblemList(db *sql.DB, page int) ([]models.Problem, error) {
	const sqlStr = `
	SELECT id, title, author, statement, created_at
	FROM problems
	LIMIT ? OFFSET ?;
	`

	rows, err := db.Query(sqlStr, ProblemNumPerPage, (page-1)*ProblemNumPerPage)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}
	defer rows.Close()

	problemList := make([]models.Problem, 0)

	for rows.Next() {
		var problem models.Problem
		err := rows.Scan(&problem.ID, &problem.Title, &problem.Author, &problem.Statement, &problem.CreatedAt)
		if err != nil {
			err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
			return nil, err
		}
		problemList = append(problemList, problem)
	}

	return problemList, nil
}

func SelectProblem(db *sql.DB, ID int) (models.Problem, error) {
	const sqlStr = `
	SELECT title, statement, answer, author, reference, reference_url, created_at
	FROM problems
	WHERE id = ?;
	`

	var problem models.Problem

	row := db.QueryRow(sqlStr, ID)
	if err := row.Err(); err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.Problem{}, err
	}

	err := row.Scan(&problem.Title, &problem.Statement, &problem.Answer, &problem.Author, &problem.Reference, &problem.ReferenceURL)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.Problem{}, err
	}

	return problem, nil
}

func InsertProblem(db *sql.DB, problem models.Problem) (models.Problem, error) {
	const sqlStr = `
	insert into problems (title, problem_statement, answer, created_at) values
	(?, ?, ?, now());
	`

	var newProblem models.Problem
	newProblem.Title, newProblem.ProblemStatement, newProblem.ProblemAnswer = problem.Title, problem.ProblemStatement, problem.ProblemAnswer

	result, err := db.Exec(sqlStr, problem.Title, problem.ProblemStatement, problem.ProblemAnswer)
	if err != nil {
		err = myerrors.InsertDataFailed.Wrap(err, "failed to insert data")
		return models.Problem{}, err
	}

	id, _ := result.LastInsertId()
	newProblem.ID = int(id)

	return newProblem, nil
}
