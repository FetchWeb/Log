package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// LogDebug determins whether debug logs should be output or not.
var LogDebug bool

// LogOutput is the file logs are written to.
var LogOutput *os.File

var logger *log.Logger
var logsDirectory string
var currentDay int

// Startup initialises the package.
func Startup(logDebug bool, logOutput *os.File) error {
	LogDebug = logDebug
	LogOutput = logOutput

	var err error
	logsDirectory, err = filepath.Abs(filepath.Dir(logOutput.Name()))
	if err != nil {
		return err
	}

	currentDay = time.Now().Day()

	logger = &log.Logger{}
	logger.SetPrefix("[MYTRAINR_API] [" + time.Now().Format("02-Jan-2006 15:04:05") + "] ")
	logger.SetOutput(LogOutput)

	return nil
}

// Info writes an info log.
func Info(message string) {
	writeLog(Data{
		Level:   "[INFO ] ",
		Message: "[" + message + "]",
	})
}

// Infof writes an info log with arguments.
func Infof(message string, args ...interface{}) {
	writeLog(Data{
		Level:   "[INFO ] ",
		Message: fmt.Sprintf("["+message+"]", args...),
	})
}

// Error writes an error log.
func Error(message string) {
	writeLog(Data{
		Level:   "[ERROR] ",
		Message: "[" + message + "]",
	})
}

// Errorf writes an error log with arguments.
func Errorf(message string, args ...interface{}) {
	writeLog(Data{
		Level:   "[ERROR] ",
		Message: fmt.Sprintf("["+message+"]", args...),
	})
}

// Debug writes a debug log.
func Debug(message string) {
	if !LogDebug {
		return
	}

	writeLog(Data{
		Level:   "[DEBUG] ",
		Message: "[" + message + "]",
	})
}

// Debugf writes a debug log with arguments.
func Debugf(message string, args ...interface{}) {
	if !LogDebug {
		return
	}

	writeLog(Data{
		Level:   "[DEBUG] ",
		Message: fmt.Sprintf("["+message+"]", args...),
	})
}

func writeLog(data Data) {
	logger.Println(data.Level + data.Message)

	if currentDay != time.Now().Day() && LogOutput != os.Stdout {
		oldLogFile := LogOutput
		defer oldLogFile.Close()

		logFile, _ := os.OpenFile(logsDirectory+"/"+time.Now().Format("2006-01-02.log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		logger.SetOutput(logFile)

		currentDay = time.Now().Day()
	}
}
