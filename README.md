# opinlog

## Overview

A very highly opinionated logger that is built on top of the [logrus](https://github.com/sirupsen/logrus) library. 

This wrapper over logrus was created to add a few features that looked useful to include on top of structured logging.
This wrapper also makes extensive use of the context object.   

For example, you might want to have the function name (along with the stack) together with the logs.
This library allows that to be done easily by using the context object.

Another example would be to mask sensitive information (such as emails) in your logs. 
You can achieve that by passing a transformation function in the key-value pair of the structured log. 

Perhaps the most opinionated piece here would be using a `Message` object instead of strings in the logs. 
You have perhaps seen a log at one line like `log.Error("unmarshal error")` 
and another line elsewhere which does `log.Error("error unmarshalling)`. 
These basically have the same meaning but it makes the logs inconsistent. 
By having a `Message` object, the idea is to force devs to keep a list of `Message` constants somewhere that can be reused. 

As mentioned, this is a highly opinionated logger that was created for personal use. 

## APIs
### Context APIs
Context APIs are used to get/update the logger object from the context object. 

`NewFromContext(ctx)`: Returns a new instance of the logger along with the updated context. 
Pass this context throughout the functions to access the same instance of the logger. 

`GetFromContext(ctx)`: Returns the existing logger stored in the context. 
Creates a new instance if an existing logger does not exist. 

`AppendFromContext(ctx, funcName)`: Updates the context and returns a new instance of the logger. 
The new logger will be updated with the function name that is passed in.

`StoreInContext(ctx, logger)`: Updates the context with the logger object.  

### Message APIs
Message is the main string that will be logged
The idea is to force devs to create a list of `Message`s that can be reused. 
(Again, this is highly opinionated.)

`NewMessage(string)`: Returns a new message object 

### Field APIs
Fields are the key-value pairs for storing structured logs.

`NewField(key, value, transforms...)`: Creates a new field with the key and value pair.
The transformation function which accepts a string as input and returns a string as output. 
These functions will be run sequentially in the order that it was passed in. 
Common use cases of the transformation functions is to obfuscate sensitive information from the logs.   

### Log APIs
These are the levels that are exposed by this wrapper. 

`Trace(Message, fields...)`: Outputs the log at trace level

`Debug(Message, fields...)`: Outputs the log at debug level

`Info(Message, fields...)`: Outputs the log at info level

`Warn(Message, fields...)`: Outputs the log at warn level

`Error(Message, fields...)`: Outputs the log at error level

`StoreFields(fields...)`: Stores the key-value fields in the logger for passing around

### Other APIs
These are other APIs that could be useful

`ChangeFunctionStackSeparator(separator)`: This changes the stack separator character.
Default separator character is `.`

`ChangeFunctionStackKey(key)`: This changes the stack function key in the logs.
Default key is `stack`

## Examples

* To get a logger and log some information
```
var (
    logInfoMessage = opinlog.NewMessage("log this message")
)

func SomeFunc(ctx context.Context, param string) {
    _, log := opinlog.NewFromContext(ctx)
    log.Info(logInfoMessage, opinlog.NewField("param", param))
    // will log somthing like "msg=log this message, param=<value>"
} 
```

* To use the stack tracing
```
var (
    logUpperFunction = opinlog.NewMessage("upper function msg")
    logLowerFunction = opinlog.NewMessage("lower function msg")
)

func UpperFunction(ctx context.Context, param string) {
    ctx, log := opinlog.AppendFromContext(ctx, "upper")
    log.Info(logUpperFunction, opinlog.NewField("param", param))
    // will log somthing like "stack=upper, msg=upper function msg, param=<value>"

    lowerFunction(ctx)
}

func lowerFunction(ctx context.Context) {
    log := opinlog.AppendFromContext(ctx, "lower")
    log.Info(logLowerFunction)
    // will log somthing like "stack=upper.lower, msg=lower function msg"
}
```

* To store key-value pairs for convenience
```
func someFunction(ctx context.Context) {
    log := opinlog.GetFromContext(ctx)
    log.StoreFields(opinlog.NewField("key", "value"))

    log.Info(opinlog.NewMessage("info 1"))
    // will log somthing like "msg=info 1, key="value""

    log.Info(opinlog.NewMessage("info 2"))
    // will log somthing like "msg=info 2, key="value""

    // optional to update the context to pass this down the call stack
    ctx = StoreInContext(ctx, log)
}
```
