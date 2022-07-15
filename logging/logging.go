package logging

import (
	"github.com/nmnjn/pronto/config"
	"github.com/zerodha/logf"
)

var logger logf.Logger = logf.New(logf.Opts{})

func SetupLogger(config config.LoggerConfig) {
	level, err := logf.LevelFromString(config.Level)

	if err != nil {
		level = logf.InfoLevel
	}

	logger = logf.New(logf.Opts{
		EnableColor:  config.EnableColor,
		Level:        level,
		EnableCaller: config.EnableCaller,
	})
}

func GetLogger() logf.Logger {
	return logger
}
