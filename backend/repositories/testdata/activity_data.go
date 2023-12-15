package testdata

import (
	"umigame-api/models"
)

var GetActivities_Basic = []models.Activity{
	{
		UserID:        1,
		Solved:        []int{1, 2, 3},
		LikedProblems: []int{1, 2, 3},
	},
}

var AddSolved_Basic = []models.Activity{
	{
		UserID:        1,
		Solved:        []int{1, 2, 3, 4},
		LikedProblems: []int{1, 2, 3},
	},
}

var AddLiked_Basic = []models.Activity{
	{
		UserID:        1,
		Solved:        []int{1, 2, 3},
		LikedProblems: []int{1, 2, 3, 4},
	},
}

var DeleteLiked_Basic = []models.Activity{
	{
		UserID:        1,
		Solved:        []int{1, 2, 3},
		LikedProblems: []int{1, 2},
	},
}
