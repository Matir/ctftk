package parser

import (
	"fmt"
)

type ConfigError struct {
	ErrMsg string
}

func (ce ConfigError) Error() string {
	return ce.ErrMsg
}

func NewConfigError(msg string, args ...interface{}) ConfigError {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return ConfigError{fmt.Sprintf("ConfigError: %s", msg)}
}
