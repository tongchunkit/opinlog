# opinlog [WIP]

A very highly opinionated logger that is built on top of the [logrus](https://github.com/sirupsen/logrus) library. 

This wrapper over logrus was created to add a few features that looked useful to include on top of structured logging.  

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

## APIs (Still WIP)
`GetFromContext(ctx)`: Returns an instance of the logger

`AppendFromContext(ctx, name)`: Returns an instance of the logger with a trace

`Persist(ctx, fields...)`: Persists the fields in the logger
 
`NewField(key, value, transforms...)`: Creates a new field with the key and value pair

`Trace(Message, fields...)`: Outputs the log at trace level

`Debug(Message, fields...)`: Outputs the log at debug level

`Info(Message, fields...)`: Outputs the log at info level

`Warn(Message, fields...)`: Opinionated warning message

`Error(Message, fields...)`: Opinionated error message
