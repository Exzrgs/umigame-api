package repositories_test

// import (
// 	"testing"
// 	"umigame-api/models"
// )

// func TestGetActivities(t *testing.T) {
// 	tests := []struct {
// 		title    string
// 		userID   int
// 		expected models.Activity
// 	}{
// 		{
// 			title:    "basic",
// 			userID:   testdata.GetActivities_Basic[0].UserID,
// 			expected: testdata.GetActivities_Basic[0],
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.title, func(t *testing.T) {
// 			got, err := repositories.GetActivities(db, test.userID)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			if got.Solved != test.expected.Solved {
// 				t.Errorf("solved problem list is expected %+v but got %+v\n", test.expected.Solved, got.Solved)
// 			}
// 			if got.LikedProblems != test.expected.LikedProblems {
// 				t.Errorf("liked problem list is expected %+v but got %+v\n", test.expected.LikedProblems, got.LikedProblems)
// 			}
// 		})
// 	}
// }

// func TestAddSolved_Basic_OK(t *testing.T) {
// 	tests := []struct {
// 		title    string
// 		userID   int
// 		expected models.Activity
// 	}{
// 		{
// 			title:    "basic",
// 			userID:   testdata.AddSolved_Basic[0].UserID,
// 			expected: testdata.AddSolved_Basic[0].Solved,
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

// func TestAddLikedProblem_Basic_OK(t *testing.T) {
// 	tests := []struct {
// 		title    string
// 		userID   int
// 		expected models.Activity
// 	}{
// 		{
// 			title:    "basic",
// 			userID:   testdata.AddLikedProblem_Basic[0].UserID,
// 			expected: testdata.AddLikedProblem_Basic[0].LikedProblems,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.title, func(t *testing.T) {
// 			if err := repositories.AddLikedProblem(db, test.userID); err != nil {
// 				t.Fatal(err)
// 			}

// 			got, err := repositories.GetActivities(db, test.userID)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			if got.LikedProblems != test.expected {
// 				t.Errorf("solved liked problem list is expected %+v but got %+v\n", test.expected, got.LikedProblems)
// 			}
// 		})
// 	}
// }

// // func TestDeleteLiked_Basic_OK() {}
