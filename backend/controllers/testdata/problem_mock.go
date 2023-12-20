package testdata

import (
	"umigame-api/models"
)

func (*serviceMock) GetProblemListService(page int) ([]models.ProblemOutline, error) {
	return problemOutlineTestData, nil
}

func (*serviceMock) GetProblemDetailService(id int) (models.ProblemDetail, error) {
	return problemDetailTestData[0], nil
}

func (*serviceMock) PostProblemService(problem models.ProblemDetail) (models.ProblemDetail, error) {
	return problemDetailTestData[0], nil
}
