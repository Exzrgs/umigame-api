package repositories_test

import (
	"testing"
	"umigame-api/models"
	"umigame-api/repositories"
	"umigame-api/repositories/testdata"
)

func TestSelectActivityList(t *testing.T) {
	tests := []struct {
		name       string
		userID     int
		problemIDs []int
		expected   []models.Activity
	}{
		{
			name:   "basic",
			userID: testdata.SelectActivityList_Basic[0].UserID,
			problemIDs: []int{
				testdata.SelectActivityList_Basic[0].ProblemID,
				testdata.SelectActivityList_Basic[1].ProblemID,
				testdata.SelectActivityList_Basic[2].ProblemID,
			},
			expected: []models.Activity{
				testdata.SelectActivityList_Basic[0],
				testdata.SelectActivityList_Basic[1],
				testdata.SelectActivityList_Basic[2],
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gots, err := repositories.SelectActivityList(db, tt.userID, tt.problemIDs)
			if err != nil {
				t.Fatal(err)
			}

			if len(gots) != len(tt.expected) {
				t.Errorf("activity list is expected %+v but got %+v\n", tt.expected, gots)
			}

			for i, got := range gots {
				if got.IsSolved != tt.expected[i].IsSolved {
					t.Errorf("solved problem list is expected %+v but got %+v\n", tt.expected[i].IsSolved, got.IsSolved)
				}
				if got.IsLiked != tt.expected[i].IsLiked {
					t.Errorf("liked problem list is expected %+v but got %+v\n", tt.expected[i].IsLiked, got.IsLiked)
				}
			}
		})
	}
}

// func TestAddSolved_Basic_OK(t *testing.T) {
// 	tests := []struct {
// 		title     string
// 		userID    int
// 		problemID int
// 		expected  models.Activity
// 	}{
// 		{
// 			title:     "basic",
// 			userID:    testdata.AddSolved_Basic[0].UserID,
// 			problemID: testdata.AddSolved_Basic[0].ProblemID,
// 			expected:  testdata.AddSolved_Basic[0].Solved,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.title, func(t *testing.T) {
// 			if err := repositories.AddSolved(db, test.userID); err != nil {
// 				t.Fatal(err)
// 			}

// 			got, err := repositories.GetActivities(db, test.userID)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			if got.Solved != test.expected {
// 				t.Errorf("solved problem list is expected %+v but got %+v\n", test.expected, got.Solved)
// 			}
// 		})
// 	}
// }

func TestChangeLiked_Basic_OK(t *testing.T) {
	tests := []struct {
		name     string
		userID   int
		activity models.Activity
		expected bool
	}{
		{
			name:     "basic",
			userID:   testdata.ChangeLiked_Basic[0].UserID,
			activity: testdata.ChangeLiked_Basic[0],
			expected: testdata.ChangeLiked_Basic[0].IsLiked,
		},
		{
			name:     "duplicate",
			userID:   testdata.ChangeLiked_Basic[1].UserID,
			activity: testdata.ChangeLiked_Basic[1],
			expected: testdata.ChangeLiked_Basic[1].IsLiked,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repositories.ChangeLiked(db, tt.userID, tt.activity); err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectActivity(db, tt.userID, tt.activity.ProblemID)
			if err != nil {
				t.Fatal(err)
			}

			if got.IsLiked != tt.expected {
				t.Errorf("solved liked problem list is expected %+v but got %+v\n", tt.expected, got.IsLiked)
			}
		})
	}
}
