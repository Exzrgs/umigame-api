package services

import "errors"

var (
	NoData = errors.New("get 0 record from db query")
	NotActivate = errors.New("activate flag is false")
)