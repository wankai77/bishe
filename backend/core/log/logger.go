package log

import (
	"backend/config/config"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var logToFile *logrus.Logger

var loggerFile string

func setLogFile(file string) {
	loggerFile = file
}

func Init() {
	// setLogFile(filepath.Join(config.Config.Log.Path, config.Config.Log.Name))
	loggerFile = filepath.Join(config.Config.Log.Path, config.Config.Log.Name)
	logFile() // Ensure logToFile is initialized
}

func Log() *logrus.Logger {
	if config.Config.Log.Model == "file" {
		return logFile()
	} else {
		if log == nil {
			log = logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2024-04-01 15:34"}
			log.SetLevel(logrus.DebugLevel)
		}
	}
	return log
}

func logFile() *logrus.Logger {
	if logToFile == nil {
		logToFile = logrus.New()
		logToFile.SetLevel(logrus.DebugLevel)

		logWriter, _ := rotatelogs.New(
			loggerFile+"_%Y%m%d.log",
			rotatelogs.WithMaxAge(30*24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		lfshook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2024-4-1 15:44",
		})
		logToFile.AddHook(lfshook)
	}
	return logToFile
}
