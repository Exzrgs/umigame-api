package services

import (
	"umigame-api/models"
	"umigame-api/myerrors"
	"umigame-api/repositories"
)

/*
一個も問題が入手できないときのエラー処理をどうするか
ここでリストのサイズ取得して処理しよう
*/
func (s *Service) GetProblemListService(page int) ([]models.ProblemOutline, error) {
	problemList, err := repositories.SelectProblemList(s.db, page)
	if err != nil {
		return nil, err
	}

	if len(problemList) == 0 {
		err := myerrors.NoData.Wrap(NoData, "0 problem found")
		return nil, err
	}

	return problemList, nil
}

func (s *Service) GetProblemDetailService(id int) (models.ProblemDetail, error) {
	problem, err := repositories.SelectProblemDetail(s.db, id)
	if err != nil {
		return models.ProblemDetail{}, err
	}

	return problem, nil
}

func (s *Service) PostProblemService(problem models.ProblemDetail) (models.ProblemDetail, error) {
	// この変数名は要検討
	newProblem, err := repositories.InsertProblem(s.db, problem)
	if err != nil {
		return models.ProblemDetail{}, err
	}

	return newProblem, nil
}
