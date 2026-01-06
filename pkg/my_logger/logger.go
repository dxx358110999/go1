package my_logger

type MyLoggerIF interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}
