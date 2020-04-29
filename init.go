package opinlog

func init() {
	// set defaults on package init
	SetLogFormat(FormatAsText)
	SetLevel(Debug)
}
