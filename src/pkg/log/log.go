package log

import (
	"errors"
	"io"
	"os"
	"otp/src/pkg/config"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var (
	logOnce sync.Once
	log     *Logger
)

func GetLoggerInstance() *Logger {
	logOnce.Do(func() {
		log = createLogger()
	})
	return log
}

func createLogger() *Logger {
	appConfig := config.GetAppConfigInstance()
	logger := &Logger{stdoutInit()}
	err := logger.setLevel(appConfig.Log.Level)
	if err != nil {
		_ = logger.setLevel("info")
	}
	return logger
}

func (l *Logger) setLevel(lvl string) error {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		err = errors.New("failed to parse level")
		return err
	}
	l.Logger.Level = level
	return nil
}

func stdoutInit() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{})
	var logWriter io.Writer = os.Stdout
	logger.SetOutput(logWriter)
	logger.SetNoLock()
	return logger
}
