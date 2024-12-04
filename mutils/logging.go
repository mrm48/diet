package mutils

import (
	"fmt"
	"log"
)

// LogConnectionError Logging should be easily filtered, split type of message from the message
// using : separator
// This is a connection error log message format for the log file
func LogConnectionError(err error) {

	slog := fmt.Sprintf("Database Connection Error: Could not connect to the database: %v", err)
	logerr := log.Output(1, slog)

	if logerr != nil {
		fmt.Sprintf("Log error: %v", logerr)
	}

}

// LogApplicationError This is for application errors, errors where there are no upstream issues
func LogApplicationError(ltype string, message string, err error) {

	slog := fmt.Sprintf("%v: %v : %v", ltype, message, err)
	logerr := log.Output(2, slog)

	if logerr != nil {
		fmt.Sprintf("Log error: %v", logerr)
	}

}

// LogMessage This is a standard log message for application messages
func LogMessage(ltype string, message string) {

	slog := fmt.Sprintf("%v: %v", ltype, message)
	logerr := log.Output(1, slog)

	if logerr != nil {
		fmt.Sprintf("Log error: %v", logerr)
	}

}
