package utils

import (
	"log"
	"os"
	"time"
)

func LogFile(logType string, filename string, prefix string, message string, query string) {
	currentTime := time.Now()

	logFile := "logs/" + logType + "/" + currentTime.Format("01-02-2006") + filename + ".log"

	f, err := os.OpenFile(logFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, prefix, log.LstdFlags)
	query = "Query : " + query
	logger.Println(message, query)
}
