package mutils

import (
	"fmt"
	"log"
)

// LogConnectionError to the app log when a connection cannot be established with the database
func LogConnectionError(err error) {

	slog := fmt.Sprintf("Database Connection Error: Could not connect to the database: %v", err)
	logerr := log.Output(1, slog)

    LogError(logerr)

}

// LogApplicationError to the app log when the application encounters an invalid value or state
func LogApplicationError(ltype string, message string, err error) {

	slog := fmt.Sprintf("%v: %v : %v", ltype, message, err)
	logerr := log.Output(2, slog)

    LogError(logerr)

}

// LogMessage to the app log for normal processing messages
func LogMessage(ltype string, message string) {

	slog := fmt.Sprintf("%v: %v", ltype, message)
	logerr := log.Output(1, slog)

    LogError(logerr)

}

// LogError to the app log when the application gets to a state where it cannot recover
func LogError(logerr error) {
	if logerr != nil {
		res := fmt.Sprintf("Log error: %v", logerr)
		log.Output(1, res)
	}
}
