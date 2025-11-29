package logger

import (
	"log"
	"time"
)

func Info(msg string) {
	writeLog("INFO", msg, nil)
}

func Error(msg string, err error) {
	writeLog("ERROR", msg, err)
}

func writeLog(level, msg string, err error) {
	now := time.Now().Unix()

	if err != nil {
		log.Printf("[%s] %d | %s: %v", level, now, msg, err)
	} else {
		log.Printf("[%s] %d | %s", level, now, msg)
	}
}
