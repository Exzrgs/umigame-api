package testdata

import (
	"umigame-api/models"
)

var ProblemTestData = []models.ProblemDetail{
	{
		ID:               1,
		Title:            "test1",
		ProblemStatement: "赤くて甘い果物は?",
		ProblemAnswer:    "りんご",
	},
	{
		ID:               2,
		Title:            "test2",
		ProblemStatement: "黄色で酸っぱい果物は?",
		ProblemAnswer:    "レモン",
	},
	{
		ID:               3,
		Title:            "test3",
		ProblemStatement: "紫で甘い果物は?",
		ProblemAnswer:    "ぶどう",
	},
}
