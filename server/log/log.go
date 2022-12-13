package log

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	panicLogger   *log.Logger
}

func New() *Logger {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Error: Log file could not be created")
	}

	var logger Logger

	logger.infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.panicLogger = log.New(file, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &logger
}

func (logger *Logger) Info(info string, a ...any) {
	logger.infoLogger.Output(2, fmt.Sprintf(info, a...))
}

func (logger *Logger) Warning(warning string, a ...any) {
	logger.warningLogger.Output(2, fmt.Sprintf(warning, a...))
}

func (logger *Logger) Error(err string, a ...any) {
	logger.errorLogger.Output(2, fmt.Sprintf(err, a...))
}

func (logger *Logger) Panic(err string, a ...any) {
	logger.panicLogger.Output(2, fmt.Sprintf(err, a...))
	panic(fmt.Sprintf(err, a...))
}
