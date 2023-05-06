package app

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

var Logger *zap.Logger

// initLog init log settings
func initLogger() {
	// newGenericLogger()
	newRotatedLogger()
	Logger.Info("logger construction successful")
}

func newRotatedLogger() {
	var cfg zap.Config
	conf := Config.StringMap("log")
	logFile := conf["logFile"]
	// errFile := conf["errFile"]

	// replace
	logFile = strings.NewReplacer(
		"{date}", LocTime().Format("20060102"),
		"{hostname}", Hostname,
	).Replace(logFile)

	fmt.Printf("logger - file=%s\n", logFile)

	// create configuration
	if Debug {
		// cfg = zap.NewDevelopmentConfig()
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cfg.Development = true
		cfg.OutputPaths = []string{"stdout", logFile}
		// cfg.ErrorOutputPaths = []string{"stderr", errFile}
	} else {
		cfg = zap.NewProductionConfig()
		cfg.OutputPaths = []string{logFile}
		// cfg.ErrorOutputPaths = []string{errFile}
	}

	// lumberjack.Logger is already safe for concurrent use, so we don't need to lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: logFile,
		MaxSize:  10, // megabytes
		// MaxBackups: 5,
		// MaxAge: 30 // 30 days
	})

	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		w,
		cfg.Level,
	)

	// init some defined fields to log
	Logger = zap.New(core).With(zap.String("hostname", Hostname))
}
