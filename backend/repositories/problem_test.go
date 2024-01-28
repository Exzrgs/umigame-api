package repositories_test

import (
	"testing"

	"umigame-api/models"
	"umigame-api/repositories"
	"umigame-api/repositories/testdata"
)

/*
すべての問題を取得する

必要な情報
problems(id, title, author, created_at, statement)
*/
func TestSelectProblemList_OK(t *testing.T) {
	tests := []struct {
		name     string
		page     int
		expected []models.Problem
	}{
		{
			name:     "basic",
			page:     1,
			expected: testdata.SelectProblemList_Basic,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := repositories.SelectProblemList(db, tt.page)
			if err != nil {
				t.Fatal(err)
			}

			if len(gotList) != len(tt.expected) {
				t.Errorf("got length is expected %v but got %v\n", len(tt.expected), len(gotList))
			}

			for index, got := range gotList {
				if got.ID != tt.expected[index].ID {
					t.Errorf("got problem is expected %+v but got %+v\n", tt.expected[index].ID, got.ID)
				}
				if got.Title != tt.expected[index].Title {
					t.Errorf("got problem is expected %+v but got %+v\n", tt.expected[index].Title, got.Title)
				}
				if got.Author != tt.expected[index].Author {
					t.Errorf("got problem is expected %+v but got %+v\n", tt.expected[index].Author, got.Author)
				}
				if got.Statement != tt.expected[index].Statement {
					t.Errorf("got problem is expected %+v but got %+v\n", tt.expected[index].Statement, got.Statement)
				}
				if *got.CreatedAt != *tt.expected[index].CreatedAt {
					t.Errorf("got problem is expected %+v but got %+v\n", tt.expected[index].CreatedAt, got.CreatedAt)
				}
			}
		})
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func Test_SelectProblem_OK(t *testing.T) {
	tests := []struct {
		name     string
		ID       int
		expected models.Problem
	}{
		{
			name:     "basic",
			ID:       testdata.SelectProblem_Basic[0].ID,
			expected: testdata.SelectProblem_Basic[0],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repositories.SelectProblem(db, tt.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != tt.expected.ID {
				t.Errorf("got problem is expected %+v but got %+v\n", tt.expected.ID, got.ID)
			}
			if got.Title != tt.expected.Title {
				t.Errorf("got problem is expected %+v but got %+v\n", tt.expected.Title, got.Title)
			}
			if got.Statement != tt.expected.Statement {
				t.Errorf("got problem is expected %+v but got %+v\n", tt.expected.Statement, got.Statement)
			}
			if got.Answer != tt.expected.Answer {
				t.Errorf("got problem is expected %+v but got %+v\n", tt.expected.Answer, got.Answer)
			}
			if got.Author != tt.expected.Author {
				t.Errorf("got problem is expected %+v but got %+v\n", tt.expected.Author, got.Author)
			}
			if got.Reference != tt.expected.Reference {
				t.Errorf("got problem is expected %+v but got %+v\n", tt.expected.Reference, got.Reference)
			}
			if got.ReferenceURL != tt.expected.ReferenceURL {
				t.Errorf("got problem is expected %+v but got %+v\n", tt.expected.ReferenceURL, got.ReferenceURL)
			}
			if *got.CreatedAt != *tt.expected.CreatedAt {
				t.Errorf("got problem is expected %+v but got %+v\n", tt.expected.CreatedAt, got.CreatedAt)
			}
		})
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////

// /*
// 問題を挿入する

// 必要な情報
// problem(title, statement, answer, author, reference, reference_url)

// チェック項目
// ・問題が挿入できるか
// */
func TestInsertProblem_Basic_OK(t *testing.T) {
	tests := []struct {
		name     string
		problem  models.Problem
		expected models.Problem
	}{
		{
			name:     "basic",
			problem:  testdata.InsertProblem_Basic[0],
			expected: testdata.InsertProblem_Basic[0],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, tt.problem)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectProblem(db, problem.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.Title != tt.expected.Title {
				t.Errorf("inserted problem is expected %+v but got %+v\n", tt.expected.Title, got.Title)
			}
			if got.Statement != tt.expected.Statement {
				t.Errorf("inserted problem is expected %+v but got %+v\n", tt.expected.Statement, got.Statement)
			}
			if got.Answer != tt.expected.Answer {
				t.Errorf("inserted problem is expected %+v but got %+v\n", tt.expected.Answer, got.Answer)
			}
			if got.Author != tt.expected.Author {
				t.Errorf("inserted problem is expected %+v but got %+v\n", tt.expected.Author, got.Author)
			}
			if got.Reference != tt.expected.Reference {
				t.Errorf("inserted problem is expected %+v but got %+v\n", tt.expected.Reference, got.Reference)
			}
			if got.ReferenceURL != tt.expected.ReferenceURL {
				t.Errorf("inserted problem is expected %+v but got %+v\n", tt.expected.ReferenceURL, got.ReferenceURL)
			}

			t.Cleanup(func() {
				if _, err := db.Exec("DELETE FROM problems WHERE id = ?", problem.ID); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

// func TestInsertProblem_NoData_NG(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		problem  models.Problem
// 		expected string
// 	}{
// 		{
// 			name:     "no title",
// 			problem:  testdata.InsertProblem_NoData[0],
// 			expected: "cause insert error",
// 		},
// 		{
// 			name:     "no statement",
// 			problem:  testdata.InsertProblem_NoData[1],
// 			expected: "cause insert error",
// 		},
// 		{
// 			name:     "no answer",
// 			problem:  testdata.InsertProblem_NoData[2],
// 			expected: "cause insert error",
// 		},
// 		{
// 			name:     "no author",
// 			problem:  testdata.InsertProblem_NoData[3],
// 			expected: "cause insert error",
// 		},
// 		{
// 			name:     "no reference",
// 			problem:  testdata.InsertProblem_NoData[4],
// 			expected: "cause insert error",
// 		},
// 		{
// 			name:     "no reference_url",
// 			problem:  testdata.InsertProblem_NoData[5],
// 			expected: "cause insert error",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			tmp, err := repositories.InsertProblem(db, test.problem)
// 			fmt.Println(tmp)
// 			if err == nil { // "==" nil
// 				t.Errorf("expected %v but got nil", test.expected)
// 			}
// 		})
// 	}
// }

func TestInsertProblem_Char_OK(t *testing.T) {
	tests := []struct {
		name     string
		problem  models.Problem
		expected models.Problem
	}{
		{
			name:     "katakana",
			problem:  testdata.InsertProblem_Char[0],
			expected: testdata.InsertProblem_Char[0],
		},
		{
			name:     "hiragana",
			problem:  testdata.InsertProblem_Char[1],
			expected: testdata.InsertProblem_Char[1],
		},
		{
			name:     "kanji",
			problem:  testdata.InsertProblem_Char[2],
			expected: testdata.InsertProblem_Char[2],
		},
		{
			name:     "half-katakana",
			problem:  testdata.InsertProblem_Char[3],
			expected: testdata.InsertProblem_Char[3],
		},
		{
			name:     "symbol",
			problem:  testdata.InsertProblem_Char[4],
			expected: testdata.InsertProblem_Char[4],
		},
		{
			name:     "space",
			problem:  testdata.InsertProblem_Char[5],
			expected: testdata.InsertProblem_Char[5],
		},
		{
			name:     "special char",
			problem:  testdata.InsertProblem_Char[6],
			expected: testdata.InsertProblem_Char[6],
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, test.problem)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectProblem(db, problem.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.Title != test.expected.Title {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Title, got.Title)
			}

			t.Cleanup(func() {
				if _, err := db.Exec("DELETE FROM problems WHERE id = ?", problem.ID); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func TestInsertProblem_TooLong_NG(t *testing.T) {
	tests := []struct {
		name     string
		problem  models.Problem
		expected string
	}{
		{
			name:     "too long title",
			problem:  testdata.InsertProblem_TooLong[0],
			expected: "cause insert error",
		},
		{
			name:     "too long author",
			problem:  testdata.InsertProblem_TooLong[1],
			expected: "cause insert error",
		},
		{
			name:     "too long reference",
			problem:  testdata.InsertProblem_TooLong[2],
			expected: "cause insert error",
		},
		{
			name:     "too long reference_url",
			problem:  testdata.InsertProblem_TooLong[3],
			expected: "cause insert error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, test.problem)
			if err == nil { // "==" nil
				t.Errorf("expected %v but got nil", test.expected)
			} else {
				t.Cleanup(func() {
					if _, err := db.Exec("DELETE FROM problems WHERE id = ?", problem.ID); err != nil {
						t.Fatal(err)
					}
				})
			}
		})
	}
}

func TestInsertProblem_Long_OK(t *testing.T) {
	tests := []struct {
		name     string
		problem  models.Problem
		expected models.Problem
	}{
		{
			name:     "long statement",
			problem:  testdata.InsertProblem_Long[0],
			expected: testdata.InsertProblem_Long[0],
		},
		{
			name:     "long answer",
			problem:  testdata.InsertProblem_Long[1],
			expected: testdata.InsertProblem_Long[1],
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, test.problem)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectProblem(db, problem.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.Statement != test.expected.Statement {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Statement, got.Statement)
			}
			if got.Answer != test.expected.Answer {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Answer, got.Answer)
			}

			t.Cleanup(func() {
				if _, err := db.Exec("DELETE FROM problems WHERE id = ?", problem.ID); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
