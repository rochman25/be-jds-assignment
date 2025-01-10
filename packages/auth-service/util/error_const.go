package util

import "errors"

var (
	NIKALREADYEXIST = errors.New("NIK Already Exist")
	NIKNOTFOUND     = errors.New("NIK not found")
	PASSWORDWRONG   = errors.New("Wrong Password")
)
