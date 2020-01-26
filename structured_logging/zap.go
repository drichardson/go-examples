package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	prod, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	example := zap.NewExample()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zapcore.InfoLevel)
	custom := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom))

	loggers := []*zap.Logger{
		prod,
		example,
		custom,
	}

	for _, l := range loggers {
		fmt.Printf("LOGGER: %v\n", l)

		l.Debug("A debug message")

		l.Debug("Shouldn't see this message")
		l.Info("An info message with fields.",
			zap.String("field1", "string value"),
			zap.Bool("field2", true),
			zap.Int("field3", 1234))

		l.Error("An error message")

	}
}
