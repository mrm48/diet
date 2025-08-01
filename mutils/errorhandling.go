package mutils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type ErrorCode int

const (
	NotConnected ErrorCode = iota
	ApplicationError
	DatabaseError
)

var errorName = map[ErrorCode]string{
	NotConnected:		"error 101:",
	ApplicationError:	"error 201:",
	DatabaseError:		"error 301:",
}

func (en ErrorCode) String() string{
	return errorName[en]
}

// WrapError wraps an error message and then calls the appropriate logger function
func WrapError(err error, msg string, logger string) error {
	if err != nil {
		if logger == "connection" {
			LogConnectionError(err)
		} else {
			LogError(err)
		}
		return errors.New(msg)
	}
	return nil
}

// WrapServiceError will determine if an error has occurred and the error message to the response
func WrapServiceError(err error, msg string, req *gin.Context, code int) (*gin.Context, error) {
	if err != nil {
		LogError(err)
		LogApplicationError(ApplicationError.String(), msg, err)
		req.IndentedJSON(code, errors.New(msg))
		return req, errors.New(msg)
	}
	return req, nil
}
