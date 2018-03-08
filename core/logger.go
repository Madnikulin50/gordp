package core

import (
	"log"
	"os"
)

var logger *log.Logger

func InitTraceLogger() {
	logger = log.New(os.Stdout, "", 0)
}

func Info(v ...interface{}) {
	logger.Print(v)
}

func Warn(v ...interface{}) {
	logger.Print(v)
}

func Error(v ...interface{}) {
	logger.Print(v)
}