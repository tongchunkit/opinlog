package opinlog

// Message refers to the log message to be used
type Message struct{
	message string
}

func (message *Message) String() string {
	return message.message
}

// NewMessage returns a new message
func NewMessage(message string) Message {
	return Message{
		message: message,
	}
}
