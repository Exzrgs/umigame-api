package testdata

import (
	"time"

	"umigame-api/models"
)

var problemOutlineTestData = []models.ProblemOutline{
	models.ProblemOutline{
		ID: 1,
		Title: "test1",
		IsSolved: false,
		CreatedAt: time.Now(),
	},
	models.ProblemOutline{
		ID: 2,
		Title: "test2",
		IsSolved: false,
		CreatedAt: time.Now(),
	},
	models.ProblemOutline{
		ID: 3,
		Title: "test3",
		IsSolved: true,
		CreatedAt: time.Now(),
	},
	models.ProblemOutline{
		ID: 4,
		Title: "test4",
		IsSolved: false,
		CreatedAt: time.Now(),
	},
}

var problemDetailTestData = []models.ProblemDetail{
	models.ProblemDetail{
		ID: 1,
		Title: "test1",
		ProblemStatement: "testes",
		ProblemAnswer: "testestes",
		UserChat: []string{
			"BmSKIgpmQIjyRpaNjPfK",
			"XyVMhhcvom",
			"LhbhQLMriGjsZrAAUKTTOrlChZLSKktSGzDDJsyYQlzpcGCxVWQvohXpYKFQXhJhjfULOSVMARDCFUOyKRvgdcPIeVOHGYCsDfgo",
			"eUllcNJRzYrkUYCWoDEzlmuXJTjcMETTnNcuqdrBPiwsTowaOKFkYRGIXditwqfZUOwbsNBaZoPGatmlFUAAkgNSdYIPYrlvYqig",
		},
		GPTChat: []string{
			"yes",
			"no",
			"don't know",
			"yes",
		},
		IsSolved: false,
		CreatedAt: time.Now(),
	},
}