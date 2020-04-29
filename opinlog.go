package opinlog

// OpinLog is the struct that implements the interface
type OpinLog struct {
	stack []string
}

func (log *OpinLog) append(name string) {
	log.stack = append(log.stack, name)
}

func (log *OpinLog) Debug(message Message, fields ...Field) {
	convertToEntry(log.stack, fields...).Debug(message.String())
}
