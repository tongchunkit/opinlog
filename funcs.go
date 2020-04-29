package opinlog

import "context"

// NewFromContext creates a new logger and stores it in the context
// to get this same logger later, you need to use the returned context
func NewFromContext(ctx context.Context) (context.Context, ILog) {
	log := &OpinLog{}
	newCtx := context.WithValue(ctx, loggerKey, log)
	return newCtx, log
}

// FromContext gets the existing logger from the context
// returns new logger if it does not exist
func GetFromContext(ctx context.Context) ILog {
	return getOpinLogFromContext(ctx)
}

// FromContextAppend gets the existing logger and appends the current name to the context
func FromContextAppend(ctx context.Context, name string) ILog {
	return appendToOpinLog(ctx, name)
}

func getOpinLogFromContext(ctx context.Context) *OpinLog {
	if value := ctx.Value(loggerKey); value != nil {
		if log, ok := value.(*OpinLog); ok {
			return log
		}
	}
	return &OpinLog{}
}

func appendToOpinLog(ctx context.Context, name string) *OpinLog {
	log := getOpinLogFromContext(ctx)
	log.append(name)
	return log
}
