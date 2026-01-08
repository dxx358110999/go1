package zap_ok

import (
	"dxxproject/config_prepare/app_config"
	"github.com/natefinch/lumberjack"
	"github.com/samber/do/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func NewZapLogger(injector do.Injector) (zapLogger *zap.Logger, err error) {
	/*
		初始化logger
	*/

	cfg := do.MustInvoke[*app_config.AppConfig](injector).LogConfig
	mode := do.MustInvoke[*app_config.AppConfig](injector).Mode

	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err2 := l.UnmarshalText([]byte(cfg.Level))
	if err2 != nil {
		return
	}

	//根据模式,输出日志
	var core zapcore.Core
	if mode == "debug" {
		// debug模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	logger := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(logger)
	zap.L().Info("zap初始化成功")

	zapLogger = logger
	return
}

func Provide(injector do.Injector) {
	do.Provide(injector, NewZapLogger)
}
