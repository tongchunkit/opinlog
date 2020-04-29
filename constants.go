package opinlog

var (
	separator = "."
)

// ChangeFunctionSeparator changes the separator used for the function stack
func ChangeFunctionSeparator(sep string) {
	separator = sep
}
