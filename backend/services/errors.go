package services

import "errors"

var (
	ErrNoData   = errors.New("get 0 record from db query")
	ErrNotValid = errors.New("email is not valified")
)
