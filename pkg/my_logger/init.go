package my_logger

import (
	"github.com/samber/do/v2"
	"go.uber.org/zap"
)

func NewMyLogger(injector do.Injector) (myLogger MyLoggerIF, err error) {
	zapLogger := do.MustInvoke[*zap.Logger](injector)
	myLogger = &ZapLoggerImpl{
		logger: zapLogger,
	}
	return
}
