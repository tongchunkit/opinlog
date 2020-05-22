package opinlog

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// NewField returns a new field
func NewField(key string, value interface{}, transformations ...func(input string) string) Field {
	return Field{
		key:             key,
		value:           value,
		transformations: transformations,
	}
}

// NewFields returns an array of fields
func NewFields(field Field, fields ...Field) []Field {
	_fields := []Field{field}
	_fields = append(_fields, fields...)
	return _fields
}

// Field is a key value pair for holding structured logs
type Field struct {
	key string
	value interface{}
	transformations []func(input string) string
}

func (field *Field) convertToKeyAndValue() (key string, value string) {
	value = fmt.Sprintf("%+v", field.value)

	for _, transform := range field.transformations {
		value = transform(value)
	}

	return field.key, value
}

func (field *Field) convertToEntry() *logrus.Entry {
	key, value := field.convertToKeyAndValue()
	return logrus.WithField(key, value)
}

func convertToEntry(trace []string, fields ...Field) *logrus.Entry {
	kvPairs := map[string]interface{}{}

	for _, field := range fields {
		key, value := field.convertToKeyAndValue()
		kvPairs[key] = value
	}

	if len(trace) > 0 {
		traceString := strings.Join(trace, functionStackSeparator)
		kvPairs[functionStackKey] = traceString
	}

	return logrus.WithFields(kvPairs)
}
