package logger

import (
	"log"
	"os"
)

var Log *Logger

type Logger struct {
	info *log.Logger
	warn *log.Logger
	err  *log.Logger
}

func (log *Logger) Info(v ...interface{}) {
	log.info.Println(v...)
}

func (log *Logger) Warn(v ...interface{}) {
	log.warn.Println(v...)
}

func (log *Logger) Error(v ...interface{}) {
	log.err.Println(v...)
}

func New(file *os.File) *Logger {
	flags := log.LstdFlags | log.Lshortfile

	log.SetOutput(file)

	return &Logger{
		info: log.New(log.Writer(), "INFO: ", flags),
		warn: log.New(log.Writer(), "WARN: ", flags),
		err:  log.New(log.Writer(), "ERROR: ", flags),
	}
}
