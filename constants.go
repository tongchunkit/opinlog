package opinlog

var (
	functionStackSeparator = "."
	functionStackKey = "stack"
)

// ChangeFunctionStackSeparator changes the separator used for the function stack
func ChangeFunctionStackSeparator(sep string) {
	functionStackSeparator = sep
}

// ChangeFunctionStackKey changes the key used to identify the function stack in the logs
func ChangeFunctionStackKey(key string) {
	functionStackKey = key
}
