package opinlog

// ILog ...
type ILog interface {
	Trace(message Message, fields ...Field)
	Debug(message Message, fields ...Field)
	Info(message Message, fields ...Field)
	Warn(message Message, fields ...Field)
	Error(message Message, fields ...Field)
	StoreFields(fields ...Field) ILog
}
