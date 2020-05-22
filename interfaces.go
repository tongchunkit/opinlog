package opinlog

// ILog is the interface used
type ILog interface {
	Trace(message Message, fields ...Field)
	TraceRaw(msg string, fields ...Field)
	Debug(message Message, fields ...Field)
	DebugRaw(msg string, fields ...Field)
	Info(message Message, fields ...Field)
	InfoRaw(msg string, fields ...Field)
	Warn(message Message, fields ...Field)
	WarnRaw(msg string, fields ...Field)
	Error(message Message, fields ...Field)
	ErrorRaw(msg string, fields ...Field)
	StoreFields(fields ...Field) ILog
}
