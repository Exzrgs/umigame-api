package testdata

import (
	"umigame-api/models"
)

var GetActivity_Basic = []models.Activity{
	{
		UserID:     1,
		SolvedList: []int{1, 2, 3},
		LikedList:  []int{1, 2, 3},
	},
}

var AddSolved_Basic = []models.Activity{
	{
		UserID:     1,
		SolvedList: []int{1, 2, 3, 4},
		LikedList:  []int{1, 2, 3},
	},
}

var AddLiked_Basic = []models.Activity{
	{
		UserID:     1,
		SolvedList: []int{1, 2, 3},
		LikedList:  []int{1, 2, 3, 4},
	},
}

var DeleteLiked_Basic = []models.Activity{
	{
		UserID:     1,
		SolvedList: []int{1, 2, 3},
		LikedList:  []int{1, 2},
	},
}
