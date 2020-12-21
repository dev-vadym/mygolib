package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Use *logrus.Logger

func NewLogger(path string) {

	log := logrus.New()
	log.SetFormatter(&Formatter{
		HideKeys:        false,
		NoColors:        true,
		ShowFullLevel:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FieldsOrder:     []string{"service", "handler"},
	})

	//logPath, _ := os.Getwd()
	logPath := fmt.Sprintf("%s/logs", path)
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.Mkdir(logPath, 0700)
	}

	logFileName := fmt.Sprintf("%s/app.log", logPath)
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to write log file:", logFile)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	//log.Info("Log file: ", logFileName)

	Use = log
}

func WithField(key string, value interface{}) *logrus.Entry {
	return Use.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return Use.WithFields(fields)
}

func Debug(args ...interface{}) {
	Use.Debug(args...)
}

func Info(args ...interface{}) {
	Use.Info(args...)
}

func Println(args ...interface{}) {
	Use.Println(args)
}

func Warn(args ...interface{}) {
	Use.Warn(args...)
}

func Error(args ...interface{}) {
	Use.Error(args...)
}

func Fatal(args ...interface{}) {
	Use.Fatal(args...)
}
