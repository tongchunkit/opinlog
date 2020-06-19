package opinlog

import "context"

type contextKey string

const (
	loggerKey contextKey = "_opinlog"
)

// GetFromContext gets the existing logger from the context
// returns new logger if it does not exist
func GetFromContext(ctx context.Context) ILog {
	return getOpinLogFromContext(ctx)
}

// AppendFromContext gets the existing logger and appends the current funcName to the stack
func AppendFromContext(ctx context.Context, funcName string) (context.Context, ILog) {
	log := getAndAppendToOpinLog(ctx, funcName)
	newCtx := StoreInContext(ctx, log)
	return newCtx, log
}

// StoreInContext stores the existing log setting in the context so that it can be retrieved later
func StoreInContext(ctx context.Context, log ILog) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

func getOpinLogFromContext(ctx context.Context) *OpinLog {
	if value := ctx.Value(loggerKey); value != nil {
		if log, ok := value.(*OpinLog); ok {
			return log
		}
	}
	return &OpinLog{}
}

func getAndAppendToOpinLog(ctx context.Context, name string) *OpinLog {
	log := getOpinLogFromContext(ctx).clone()
	log.append(name)
	return log
}
