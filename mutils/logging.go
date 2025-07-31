package mutils

import (
	"fmt"
	"log"
)

type LogType int

const (
	Request LogType = iota
	Debug
	ServerStartup
)

var logType = map[LogType]string{
	Request:       "Request:",
	Debug:         "Debug:",
	ServerStartup: "Server Startup:",
}

func (lt LogType) String() string {
	return logType[lt]
}

var callDepthLogged int = 2

// LogConnectionError to the app log when a connection cannot be established with the database
func LogConnectionError(err error) {

	slog := fmt.Sprintf("Database Connection Error: Could not connect to the database: %v", err)
	logerr := log.Output(callDepthLogged, slog)

	LogError(logerr)

}

// LogApplicationError to the app log when the application encounters an invalid value or state
func LogApplicationError(ltype string, message string, err error) {

	slog := fmt.Sprintf("%v: %v : %v", ltype, message, err)

	logerr := log.Output(callDepthLogged, slog)

	LogError(logerr)

}

// LogMessage to the app log for normal processing messages
func LogMessage(ltype string, message string) {

	slog := fmt.Sprintf("%v: %v", ltype, message)
	logerr := log.Output(callDepthLogged, slog)

	LogError(logerr)

}

// LogError to the app log when the application gets to a state where it cannot recover
func LogError(logerr error) {
	if logerr != nil {
		res := fmt.Sprintf("Log error: %v", logerr)
		log.Output(callDepthLogged, res)
	}
}
