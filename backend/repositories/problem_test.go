package repositories_test

import (
	"testing"
	"database/sql"
	"errors"

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
		expected []models.Problem
	}{
		{
			name:     "1",
			expected: testdata.SelectProblemList_Basic,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			gotList, err := repositories.SelectProblemList(db, test.page)
			if err != nil {
				t.Fatal(err)
			}

			for index, got := range gotList {
				if got != test.expected[index] {
					t.Errorf("problem list is expected %+v but got %+v\n", test.expected, got)
				}
			}
		})
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func Test_SelectProblem_OK(t *testing.T) {
	tests := []struct {
		title    string
		ID       int
		expected models.Problem
	}{
		{
			title:    "basic",
			ID:       testdata.SelectProblemDetail_Basic[0].ID,
			expected: testdata.SelectProblemDetail_Basic[0],
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			got, err := repositories.SelectProblemDetail(db, test.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got != test.expected {
				t.Errorf("got problem is expected %+v but got %+v\n", test.expected, got)
			}
		})
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////

/*
問題を挿入する

必要な情報
problem(title, statement, answer, author, reference, reference_url)

チェック項目
・問題が挿入できるか
*/
func TestInsertProblem_Basic_OK(t *testing.T) {
	tests := []struct {
		title    string
		problem  models.Problem
		expected models.Problem
	}{
		{
			title:    "basic",
			problem:  testdata.InsertProblem_Basic[0],
			expected: testdata.InsertProblem_Basic[0],
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, test.problem)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectProblemDetail(db, problem.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.Title != test.expected.Title {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Title, got.Title)
			}
			if got.Statement != test.expected.Statement {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Statement, got.Statement)
			}
			if got.Answer != test.expected.Answer {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Answer, got.Answer)
			}
			if got.Author != test.expected.Author {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Author, got.Author)
			}
			if got.Reference != test.expected.Reference {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Reference, got.Reference)
			}
			if got.ReferenceURL != test.expected.ReferenceURL {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.ReferenceURL, got.ReferenceURL)
			}
		})
	}
}

func TestInsertProblem_NoData_NG(t *testing.T) {
	tests := []struct {
		title    string
		problem  models.Problem
		expected string
	}{
		{
			title:    "no title",
			problem:  testdata.InsertProblem_NoData[0],
			expected: "cause error"
		},
		{
			title:    "no statement",
			problem:  testdata.InsertProblem_NoData[1],
			expected: "cause error"
		},
		{
			title:    "no answer",
			problem:  testdata.InsertProblem_NoData[2],
			expected: "cause error"
		},
		{
			title:    "no author",
			problem:  testdata.InsertProblem_NoData[3],
			expected: "cause error"
		},
		{
			title:    "no reference",
			problem:  testdata.InsertProblem_NoData[4],
			expected: "cause error"
		},
		{
			title:    "no reference_url",
			problem:  testdata.InsertProblem_NoData[5],
			expected: "cause error"
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, test.problem)
			if err != nil {
				t.Fatal(err)
			}

			if err == nil { // "==" nil
				t.Errorf("expected %v but got %v", test.expected, got)
			}
		})
	}
}

func TestInsertProblem_Char_OK(t *testing.T) {
	tests := []struct {
		title    string
		problem  models.Problem
		expected models.Problem
	}{
		{
			title:    "katakana",
			problem:  testdata.InsertProblem_Char[0],
			expected: testdata.InsertProblem_Char[0],
		},
		{
			title:    "hiragana",
			problem:  testdata.InsertProblem_Char[1],
			expected: testdata.InsertProblem_Char[1],
		},
		{
			title:    "kanji",
			problem:  testdata.InsertProblem_Char[2],
			expected: testdata.InsertProblem_Char[2],
		},
		{
			title:    "half-katakana",
			problem:  testdata.InsertProblem_Char[3],
			expected: testdata.InsertProblem_Char[3],
		},
		{
			title:    "symbol",
			problem:  testdata.InsertProblem_Char[4],
			expected: testdata.InsertProblem_Char[4],
		},
		{
			title:    "space",
			problem:  testdata.InsertProblem_Char[5],
			expected: testdata.InsertProblem_Char[5],
		},
		{
			title:    "special char",
			problem:  testdata.InsertProblem_Char[6],
			expected: testdata.InsertProblem_Char[6],
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, test.problem)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectProblemDetail(db, problem.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.Title != test.expected.Title {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Title, got.Title)
			}
		})
	}
}

func TestInsertProblem_TooLong_NG(t *testing.T) {
	tests := []struct {
		title    string
		problem  models.Problem
		expected string
	}{
		{
			title:    "too long title",
			problem:  testdata.InsertProblem_TooLong[0],
			expected: "cause error"
		},
		{
			title:    "too long author",
			problem:  testdata.InsertProblem_TooLong[1],
			expected: "cause error"
		},
		{
			title:    "too long reference",
			problem:  testdata.InsertProblem_TooLong[2],
			expected: "cause error"
		},
		{
			title:    "too long author",
			problem:  testdata.InsertProblem_TooLong[3],
			expected: "cause error"
		},
		{
			title:    "too long reference",
			problem:  testdata.InsertProblem_TooLong[4],
			expected: "cause error"
		},
		{
			title:    "too long reference_url",
			problem:  testdata.InsertProblem_TooLong[5],
			expected: "cause error"
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, test.problem)
			if err != nil {
				t.Fatal(err)
			}

			if err == nil { // "==" nil
				t.Errorf("expected %v but got %v", test.expected, got)
			}
		})
	}
}

func TestInsertProblem_Long_OK(t *testing.T) {
	tests := []struct {
		title    string
		problem  models.Problem
		expected models.Problem
	}{
		{
			title:    "long statement",
			problem:  testdata.InsertProblem_Long[0],
			expected: testdata.InsertProblem_Long[0],
		},
		{
			title:    "long answer",
			problem:  testdata.InsertProblem_Long[1],
			expected: testdata.InsertProblem_Long[1],
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, test.problem)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectProblemDetail(db, problem.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.Statement != test.expected.Statement {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Statement, got.Statement)
			}
			if got.Answer != test.expected.Answer {
				t.Errorf("inserted problem is expected %+v but got %+v\n", test.expected.Answer, got.Answer)
			}
		})
	}
}