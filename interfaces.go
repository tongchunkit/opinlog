package opinlog

// ILog ...
type ILog interface {
	Debug(message Message, fields ...Field)
	//Info(message Message)
	//Warn(message Message)
	//Error(message Message)
	//Fatal(message Message)
}
