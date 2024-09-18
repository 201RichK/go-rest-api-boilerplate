package logger

import (
	"os"
	"path/filepath"

	"github.com/aro/configs"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

type Level int

func Setup() (func(), error) {
	logger = logrus.New()

	// Set the default log level
	logger.SetLevel(logrus.DebugLevel)

	// Custom formatter for color
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		FullTimestamp: true,
	})

	var file *os.File
	output := configs.LoggerSetting.Output
	if output != "" {
		switch output {
		case "stdout":
			logger.SetOutput(os.Stdout)
		case "stderr":
			logger.SetOutput(os.Stderr)
		case "file":
			fileName := configs.LoggerSetting.Filename
			_ = os.MkdirAll(filepath.Dir(fileName), 0777)

			f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				return nil, err
			}
			logger.SetOutput(f)
			file = f
		}
	}

	return func() {
		if file != nil {
			file.Close()
		}
	}, nil
}

// Colorful log wrappers
func Debug(v ...interface{}) {
	color.New(color.FgCyan).Add(color.Bold).Print("[DEBUG] ")
	logger.Debug(v...)
}

func Debugf(s string, v ...interface{}) {
	color.New(color.FgCyan).Add(color.Bold).Print("[DEBUG] ")
	logger.Debugf(s, v...)
}

func Info(v ...interface{}) {
	color.New(color.FgGreen).Add(color.Bold).Print("[INFO] ")
	logger.Info(v...)
}

func Infof(s string, v ...interface{}) {
	color.New(color.FgGreen).Add(color.Bold).Print("[INFO] ")
	logger.Infof(s, v...)
}

func Warn(v ...interface{}) {
	color.New(color.FgYellow).Add(color.Bold).Print("[WARN] ")
	logger.Warn(v...)
}

func Warnf(s string, v ...interface{}) {
	color.New(color.FgYellow).Add(color.Bold).Print("[WARN] ")
	logger.Warnf(s, v...)
}

func Error(v ...interface{}) {
	color.New(color.FgRed).Add(color.Bold).Print("[ERROR] ")
	logger.Error(v...)
}

func Errorf(s string, v ...interface{}) {
	color.New(color.FgRed).Add(color.Bold).Print("[ERROR] ")
	logger.Errorf(s, v...)
}

func Fatal(v ...interface{}) {
	color.New(color.FgRed).Add(color.Bold).Print("[FATAL] ")
	logger.Fatal(v...)
}

func Fatalf(s string, v ...interface{}) {
	color.New(color.FgRed).Add(color.Bold).Print("[FATAL] ")
	logger.Fatalf(s, v...)
}
