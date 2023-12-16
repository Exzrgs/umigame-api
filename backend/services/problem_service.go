package services

import (
	"umigame-api/models"
	"umigame-api/myerrors"
	"umigame-api/repositories"
)

/*
pageは必要
アクティビティも取得
*/
func (s *Service) GetProblemListService(page int) (map[int]models.ProblemBase, error) {
	problemList, err := repositories.SelectProblemList(s.db, page)
	if err != nil {
		return nil, err
	}
	if len(problemList) == 0 {
		err := myerrors.NoData.Wrap(NoData, "0 problem found")
		return nil, err
	}

	problemIDs := make([]int, ProblemNumPerPage)
	for i, problem := range problemList {
		problemIDs[i] = problem.ID
	}
	activityList, err := repositories.GetActivityList(s.db, userID, problemIDs)
	if err != nil {
		return nil, err
	}

	response := make(map[int]models.ProblemBase, ProblemNumPerPage)
	for _, problem := range problemList {
		var problemBase models.ProblemBase
		problemBase.Title, problemBase.Author, problemBase.Statement, problemBase.CreatedAt = problem.Title, problem.Author, problem.Statement, problem.CreatedAt
		problemBase.IsSolved, problemBase.IsLiked = false, false
		response[problem.ID] = problemBase
	}
	for _, activity := range activityList {
		response[activity.ProblemID].IsSolved = activity.IsSolved
		response[activity.ProblemID].IsLiked = activity.IsLiked
	}

	return response, nil
}

func (s *Service) GetProblemDetailService(id int) (models.Problem, error) {
	problem, err := repositories.SelectProblemDetail(s.db, id)
	if err != nil {
		return models.Problem{}, err
	}

	return problem, nil
}

func (s *Service) PostProblemService(problem models.Problem) (models.Problem, error) {
	// この変数名は要検討
	newProblem, err := repositories.InsertProblem(s.db, problem)
	if err != nil {
		return models.Problem{}, err
	}

	return newProblem, nil
}
