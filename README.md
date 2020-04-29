# opinlog
A very highly opinionated logger

## APIs
`GetFromContext(ctx)`: Returns an instance of the logger

`AppendFromContext(ctx, name)`: Returns an instance of the logger with a trace

`Persist(ctx, fields...)`: Persists the fields in the logger
 
`NewField(key, value, transforms...)`: Creates a new field with the key and value pair

`Trace(msg, fields...)`: Outputs the log at trace level

`Debug(msg, fields...)`: Outputs the log at debug level

`Info(msg, fields...)`: Outputs the log at info level

`Warn(Message, fields...)`: Opinionated warning message

`Error(Message, fields...)`: Opinionated error message
