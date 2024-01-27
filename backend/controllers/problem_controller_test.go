package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetProblemListHandler(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		resCode int
	}{
		{
			"test1",
			"1",
			http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			url := fmt.Sprintf("%s/problem/list?page=%s", baseURL, test.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			pc.GetProblemListHandler(w, req)

			if w.Code != test.resCode {
				t.Errorf("unexpected status code: want %d but get %d\n", test.resCode, w.Code)
			}
		})
	}
}

func TestGetProblemDetailHandler(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		resCode int
	}{
		{
			"test1",
			"1",
			http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			url := fmt.Sprintf("%s/problem/detail?id=%s", baseURL, test.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			pc.GetProblemDetailHandler(w, req)

			if w.Code != test.resCode {
				t.Errorf("unexpected status code: want %d but get %d\n", test.resCode, w.Code)
				return
			}
		})
	}
}

func PostProblemHandler(t *testing.T) {
	tests := []struct {
		name    string
		resCode int
		body    string
	}{
		{
			"test1",
			http.StatusOK,
			`{
				"id": "1",
				"title": "tes",
				"problemstatement": "tes",
				"problem_answer": "tes",
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			url := fmt.Sprintf("%s/problem/add", baseURL)
			reqBody := strings.NewReader()
			req := httptest.NewRequest(http.MethodPost, url)
			pc.PostProblemHandler(w, req)

			if w.Code != test.resCode {
				t.Errorf("unexpected status code: want %d but get %d\n", test.resCode, w.Code)
				return
			}
		})
	}
}
