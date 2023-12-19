package myerrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"umigame-api/utils"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var myErr *MyError
	if !errors.As(err, &myErr) {
		myErr = &MyError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := utils.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, myErr)

	var statusCode int

	switch myErr.ErrCode {
	case NoData:
		statusCode = http.StatusNotFound
	case ReqDecodeFailed, BadParameter, BadCookie:
		statusCode = http.StatusBadRequest
	case InvalidPassword, InvalidUUID, NotActivate:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(myErr)
}
