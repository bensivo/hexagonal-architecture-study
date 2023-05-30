package logging

import "go.uber.org/zap"

var Sugar *zap.SugaredLogger

func Init() *zap.SugaredLogger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return logger.Sugar()
}
