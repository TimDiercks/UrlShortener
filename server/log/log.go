package log

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func New() Logger {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Error: Log file could not be created")
	}

	var logger Logger

	logger.infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}

func (logger *Logger) Info(info string) {
	logger.infoLogger.Output(2, info)
}

func (logger *Logger) Warning(warning string) {
	logger.warningLogger.Output(2, warning)
}

func (logger *Logger) Error(err string) {
	logger.errorLogger.Output(2, err)
}
