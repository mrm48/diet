package mutils

import (
	"fmt"
	"log"
)

func LogConnectionError(err error) {

	slog := fmt.Sprintf("Cannot connect to the database: %v", err)
	log.Output(1, slog)

}

func LogApplicationError(ltype string, message string, err error) {

	slog := fmt.Sprintf("%v: %v : %v", ltype, message, err)
	log.Output(2, slog)

}

func LogMessage(ltype string, message string) {

	slog := fmt.Sprintf("%v: %v", ltype, message)
	log.Output(1, slog)

}
