package mutils

import "errors"

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
