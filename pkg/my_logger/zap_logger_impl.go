package my_logger

import (
	"go.uber.org/zap"
)

/*
实现logger接口
*/

type ZapLoggerImpl struct {
	logger *zap.Logger
}

func (r *ZapLoggerImpl) Debug(msg string) {
	r.logger.Debug(msg)
}
func (r *ZapLoggerImpl) Info(msg string) {
	r.logger.Info(msg)

}
func (r *ZapLoggerImpl) Warn(msg string) {
	r.logger.Warn(msg)

}
func (r *ZapLoggerImpl) Error(msg string) {
	r.logger.Error(msg)

}

var _ MyLoggerIF = new(ZapLoggerImpl)
