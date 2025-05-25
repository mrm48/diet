package mutils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"net/http"
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

func WrapServiceError(err error, msg string, logger string, req *gin.Context) error {
	if err != nil {
		LogError(err)
		LogApplicationError("Application Error", msg, err)
		req.IndentedJSON(http.StatusInternalServerError, errors.New(msg))
	} 
	return errors.New(msg)
}
