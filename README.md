# opinlog

## Overview

A personal logger that is built on top of the [logrus](https://github.com/sirupsen/logrus) library. 

This wrapper over logrus was created to add a few features that looked useful to include on top of structured logging.
This wrapper also makes extensive use of the context object.   

For example, you might want to have the function name (along with the stack) together with the logs.
This library allows that to be done easily by using the context object.

Another example would be to mask sensitive information (such as emails) in your logs 
(if external logging filters cannot catch those via regular expressions). 
You can achieve that by passing a transformation function when creating a key-value pair of the object in the structured log. 

A helpful wrapper here would be the `Message` object.  
You have perhaps seen one line which does `log.Error("unmarshal error")` 
and another line elsewhere which does `log.Error("error unmarshalling object")`. 
These basically have the same meaning but it makes the logs inconsistent. 
By having a `Message` object, the idea is to force devs to keep a list of `Message` vars somewhere that can be reused. 

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
Using `Message` is a convenient way to wrap logging strings and `Fields` in a single call 
rather than using strings and `Fields` each time.  

Message is a function which returns the main logging string as well as key-value field pairs. 

Sample of how to create a `Message`:
```
func UnmarshalError(obj interface{}, err error) opinlog.Message {
  return func() (string, []opinlog.Field) {
    return "unmarshal error",
      opinlog.NewFields(opinlog.NewField("err", err), opinlog.NewField("obj", obj))
    }
  }
```

How to use a created `Message`: 
```
func yourFunction(ctx context.Context, obj interface{}) {
  logger := opinlog.GetFromContext(ctx)

  err := json.Unmarshal(...)
  if err != nil {
    logger.Error(UnmarshalError(obj, err))
  }
}
```

### Field APIs
Fields are the key-value pairs for storing structured logs.

`NewField(key, value, transforms...)`: Creates a new field with the key and value pair.
The transformation function which accepts a string as input and returns a string as output. 
These functions will be run sequentially in the order that it was passed in. 
Common use cases of the transformation functions is to obfuscate sensitive information from the logs.

`NewFields(field, ...fields)`: Returns an array of field objects   

### Log APIs
These are the levels that are exposed by this wrapper. 

`Trace(Message, fields...)`: Outputs the log at trace level

`TraceRaw(string, fields...)`: Outputs the log at trace level

`Debug(Message, fields...)`: Outputs the log at debug level

`DebugRaw(string, fields...)`: Outputs the log at debug level

`Info(Message, fields...)`: Outputs the log at info level

`InfoRaw(string, fields...)`: Outputs the log at info level

`Warn(Message, fields...)`: Outputs the log at warn level

`WarnRaw(string, fields...)`: Outputs the log at warn level

`Error(Message, fields...)`: Outputs the log at error level

`ErrorRaw(string, fields...)`: Outputs the log at error level

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
func SomeFunc(ctx context.Context, param string) {
    _, log := opinlog.NewFromContext(ctx)
    log.InfoRaw("log this message", opinlog.NewField("param", param))
    // will log somthing like "msg=log this message, param=<value>"
} 
```

* To use the stack tracing
```
func UpperFunction(ctx context.Context, param string) {
    ctx, log := opinlog.AppendFromContext(ctx, "upper")
    log.InfoRaw("upper function msg", opinlog.NewField("param", param))
    // will log somthing like "stack=upper, msg=upper function msg, param=<value>"

    lowerFunction(ctx)
}

func lowerFunction(ctx context.Context) {
    log := opinlog.AppendFromContext(ctx, "lower")
    log.InfoRaw("lower function msg")
    // will log somthing like "stack=upper.lower, msg=lower function msg"
}
```

* To store key-value pairs for convenience
```
func someFunction(ctx context.Context) {
    log := opinlog.GetFromContext(ctx)
    log.StoreFields(opinlog.NewField("key", "value"))

    log.InfoRaw("info 1")
    // will log somthing like "msg=info 1, key="value""

    log.InfoRaw("info 2")
    // will log somthing like "msg=info 2, key="value""

    // to update the context to pass this down in other functions
    ctx = StoreInContext(ctx, log)
}
```

* To use the `Message` object
```
<see above in Message APIs>
```
