package my_logger

import (
	"go.uber.org/zap"
)

/*
实现logger接口
*/

type MyLoggerZapImpl struct {
	logger *zap.Logger
}

func fieldsToZapFields(others ...LogField) []zap.Field {
	size := len(others)
	zapFields := make([]zap.Field, size)

	for i := 0; i < size; i++ {
		f := zap.Any(others[i].Key, others[i].Value)
		zapFields[0] = f
	}

	return zapFields
}

func (r *MyLoggerZapImpl) Debug(msg string, others ...LogField) {
	zapFields := fieldsToZapFields(others...)
	r.logger.Debug(msg, zapFields...)
}
func (r *MyLoggerZapImpl) Info(msg string, others ...LogField) {
	zapFields := fieldsToZapFields(others...)
	r.logger.Info(msg, zapFields...)
}
func (r *MyLoggerZapImpl) Warn(msg string, others ...LogField) {
	zapFields := fieldsToZapFields(others...)
	r.logger.Warn(msg, zapFields...)
}
func (r *MyLoggerZapImpl) Error(msg string, others ...LogField) {
	zapFields := fieldsToZapFields(others...)
	r.logger.Error(msg, zapFields...)
}

var _ MyLoggerIF = new(MyLoggerZapImpl)

func NewMyLogger(zapLogger *zap.Logger) (*MyLoggerZapImpl, error) {
	myLogger := &MyLoggerZapImpl{
		logger: zapLogger,
	}
	return myLogger, nil
}
