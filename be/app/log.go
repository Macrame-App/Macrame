package app

import (
	"log"
	"os"
)

var logFile *os.File

func MCRMLogInit() {
	var err error
	// Open or create the log file with append permissions
	logFile, err = os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err) // Exit if we can't open the log file
	}

	// Optionally set log to write to file in addition to standard log output
	log.SetOutput(logFile)
}

func MCRMLog(v ...interface{}) {
	log.Println(v...) // Logs to terminal as well
}
