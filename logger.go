package main

import (
	"log"
	"os"
	"path/filepath"
)

func initLogger() *log.Logger {
	logPath := "C:\\PriorityQueue"
	fileName := "logs.txt"
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.Mkdir(logPath, os.ModeDir)
	}

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(filepath.Join(logPath, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}
