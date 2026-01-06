package my_logger

/*
允许传入键值对作为附加信息
*/

type LogField struct {
	Key   string
	Value any
}

type MyLoggerIF interface {
	Debug(msg string, others ...LogField)
	Info(msg string, others ...LogField)
	Warn(msg string, others ...LogField)
	Error(msg string, others ...LogField)
}
