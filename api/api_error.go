package api

import (

)

type Error interface {
	ErrCode() int
}

// Error response for API calls
type apiError struct {
	Code 	int 		`json:"code,omitempty"`
	Msg 		string	`json:"msg,omitempty"`
}

func ApiError(code int, msg string) Error {
	return &apiError{
		Code: code,
		Msg: msg,
	}
}

func (err *apiError) ErrCode() int {
	return err.Code
}

const (
	ERR_BAD_REQUEST = 400
	ERR_NOT_FOUND = 404
	ERR_METHOD_NOT_ALLOWED = 405
	ERR_CONFLICT = 409
	ERR_FAILED = 500
	ERR_UNAVILABLE = 503
)

 var ErrMsgs map[int]string
 
 func init() {
 	ErrMsgs = make(map[int]string)
 	ErrMsgs[ERR_CONFLICT] = "already exists"
 	ErrMsgs[ERR_UNAVILABLE] = "service not available"
 	ErrMsgs[ERR_NOT_FOUND] = "not found"
 }
 