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
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Trace(message.String())
}

// Debug is the debug log
func (log *OpinLog) Debug(message Message, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Debug(message.String())
}

// Info is the info log
func (log *OpinLog) Info(message Message, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Info(message.String())
}

// Warn is the warn log
func (log *OpinLog) Warn(message Message, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Warn(message.String())
}

// Error is the error log
func (log *OpinLog) Error(message Message, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Error(message.String())
}

// StoreFields stores the field in the struct so that it is easy to pass around
func (log *OpinLog) StoreFields(fields ...Field) ILog {
	log.fields = append(log.fields, fields...)
	return log
}
