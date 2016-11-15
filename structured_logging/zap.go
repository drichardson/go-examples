package main

import (
	"github.com/uber-go/zap"
)

func main() {
	loggers := []zap.Logger{
		zap.New(zap.NewTextEncoder()),
		zap.New(zap.NewJSONEncoder()),
		zap.New(zap.NewJSONEncoder(zap.NoTime())),
	}

	for _, l := range loggers {

		l.SetLevel(zap.DebugLevel)

		l.Debug("A debug message")

		l.SetLevel(zap.InfoLevel)

		l.Debug("Shouldn't see this message")
		l.Info("An info message with fields.",
			zap.String("field1", "string value"),
			zap.Bool("field2", true),
			zap.Int("field3", 1234))

		l.Error("An error message")

	}
}
