package controllers_test

import (
	"testing"

	"umigame-api/controllers"
	"umigame-api/controllers/testdata"
)

var (
	baseURL = "http://localhost:8080"
	pc      *controllers.ProblemController
)

func TestMain(m *testing.M) {
	s := testdata.NewServiceMock()
	pc = controllers.NewProblemController(s)

	m.Run()
}
