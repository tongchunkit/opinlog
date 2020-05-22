package opinlog

// OpinLog is the struct that implements the interface
type OpinLog struct {
	stack []string
	fields []Field
}

func (log *OpinLog) clone() *OpinLog{
	clone := &OpinLog{}
	clone.stack = log.stack
	return clone
}

func (log *OpinLog) append(name string) {
	log.stack = append(log.stack, name)
}

// Trace is the trace log
func (log *OpinLog) Trace(message Message, fields ...Field) {
	_msg, _fields := message()
	log.TraceRaw(_msg, appendFields(_fields, fields...)...)
}

// TraceRaw is the trace log
func (log *OpinLog) TraceRaw(msg string, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Trace(msg)
}

// Debug is the debug log
func (log *OpinLog) Debug(message Message, fields ...Field) {
	_msg, _fields := message()
	log.DebugRaw(_msg, appendFields(_fields, fields...)...)
}

// DebugRaw is the debug log
func (log *OpinLog) DebugRaw(msg string, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Debug(msg)
}

// Info is the info log
func (log *OpinLog) Info(message Message, fields ...Field) {
	_msg, _fields := message()
	log.InfoRaw(_msg, appendFields(_fields, fields...)...)
}

// InfoRaw is the info log
func (log *OpinLog) InfoRaw(msg string, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Info(msg)
}

// Warn is the warn log
func (log *OpinLog) Warn(message Message, fields ...Field) {
	_msg, _fields := message()
	log.WarnRaw(_msg, appendFields(_fields, fields...)...)
}

// WarnRaw is the warn log
func (log *OpinLog) WarnRaw(msg string, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Warn(msg)
}

// Error is the error log
func (log *OpinLog) Error(message Message, fields ...Field) {
	_msg, _fields := message()
	log.ErrorRaw(_msg, appendFields(_fields, fields...)...)
}

// ErrorRaw is the error log
func (log *OpinLog) ErrorRaw(msg string, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Error(msg)
}

// StoreFields stores the field in the struct so that it is easy to pass around
func (log *OpinLog) StoreFields(fields ...Field) ILog {
	log.fields = append(log.fields, fields...)
	return log
}

func appendFields(fields []Field, moreFields ...Field) []Field {
	if fields == nil {
		fields = []Field{}
	}

	fields = append(fields, moreFields...)
	return fields
}
