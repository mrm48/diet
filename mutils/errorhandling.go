package mutils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

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

func WrapServiceError(err error, msg string, req *gin.Context, code int) (*gin.Context, error) {
	if err != nil {
		LogError(err)
		LogApplicationError("Application Error", msg, err)
		req.IndentedJSON(code, errors.New(msg))
	}
	return req, errors.New(msg)
}
