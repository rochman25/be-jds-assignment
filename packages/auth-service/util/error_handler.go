package util

import "net/http"

func GetErrorCode(err error) int {
	found := MAPERROR[err]
	return found
}

var MAPERROR = map[error]int{
	NIKALREADYEXIST: http.StatusBadRequest,
	NIKNOTFOUND:     http.StatusBadRequest,
	PASSWORDWRONG:   http.StatusBadRequest,
}
