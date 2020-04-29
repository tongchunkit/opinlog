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

func (log *OpinLog) Trace(message Message, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Trace(message.String())
}

func (log *OpinLog) Debug(message Message, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Debug(message.String())
}

func (log *OpinLog) Info(message Message, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Info(message.String())
}

func (log *OpinLog) Warn(message Message, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Warn(message.String())
}

func (log *OpinLog) Error(message Message, fields ...Field) {
	log.StoreFields(fields...)
	convertToEntry(log.stack, log.fields...).Error(message.String())
}

func (log *OpinLog) StoreFields(fields ...Field) ILog {
	log.fields = append(log.fields, fields...)
	return log
}
