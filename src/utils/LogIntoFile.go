package utils

import (
	"log"
	"os"
)

func Logger(logType, method, message string) {
	// create info log
	fileInfo, err := openLogFile("./messages.log")
	if err != nil {
		log.Fatal(err)
	}
	infoLog := log.New(fileInfo, "["+logType+"]-["+method+"] ", log.LstdFlags|log.Lmicroseconds)
	infoLog.Println(message)
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
