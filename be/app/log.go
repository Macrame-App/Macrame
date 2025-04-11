package app

import (
	"fmt"
	"log"
	"os"
)

var logFile *os.File

func MCRMLogInit() {
	var err error
	logFile, err = os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func MCRMLog(v ...interface{}) {
	log.Println(v...)
	fmt.Fprintln(logFile, v...)
}
