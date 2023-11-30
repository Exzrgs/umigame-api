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
	ProblemNumPerPage = 2
)

// 1ページ何個取得するかを決めないといけない
// 一旦20問
func SelectProblemList(db *sql.DB, page int) ([]models.ProblemOutline, error) {
	const sqlStr = `
	select id, title, is_solved, created_at
	from problems
	limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, ProblemNumPerPage, (page-1)*ProblemNumPerPage)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}
	defer rows.Close()

	problemList := make([]models.ProblemOutline, 0)

	for rows.Next() {
		var problem models.ProblemOutline
		err := rows.Scan(&problem.ID, &problem.Title, &problem.IsSolved, &problem.CreatedAt)
		if err != nil {
			err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
			return nil, err
		}
		problemList = append(problemList, problem)
	}

	return problemList, nil
}

func SelectProblemDetail(db *sql.DB, ID int) (models.ProblemDetail, error) {
	const sqlStr = `
	select id, title, problem_statement, answer
	from problems
	where id = ?;
	`

	var problem models.ProblemDetail

	row := db.QueryRow(sqlStr, ID)
	if err := row.Err(); err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.ProblemDetail{}, err
	}

	err := row.Scan(&problem.ID, &problem.Title, &problem.ProblemStatement, &problem.ProblemAnswer)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.ProblemDetail{}, err
	}

	return problem, nil
}

func InsertProblem(db *sql.DB, problem models.ProblemDetail) (models.ProblemDetail, error) {
	const sqlStr = `
	insert into problems (title, problem_statement, answer, created_at) values
	(?, ?, ?, now());
	`

	var newProblem models.ProblemDetail
	newProblem.Title, newProblem.ProblemStatement, newProblem.ProblemAnswer = problem.Title, problem.ProblemStatement, problem.ProblemAnswer

	result, err := db.Exec(sqlStr, problem.Title, problem.ProblemStatement, problem.ProblemAnswer)
	if err != nil {
		err = myerrors.InsertDataFailed.Wrap(err, "failed to insert data")
		return models.ProblemDetail{}, err
	}

	id, _ := result.LastInsertId()
	newProblem.ID = int(id)

	return newProblem, nil
}
