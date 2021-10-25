package utils

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// WrapError => use this function to wrap stack trace of errors
// when performing error handling
func WrapError(err error, message string) error {
	PushLogf(message, err.Error())
	return errors.Wrap(err, fmt.Sprintf("%s =>", message))
}

// UnwrapError => this function unwraps the wrapped errors
// returns a slice of strings, containing stack trace errors
func UnwrapError(err error) []string {
	var result []string
	splitted := strings.Split(err.Error(), "=>:")

	for _, val := range splitted {
		result = append(result, fmt.Sprintf("[%s]", strings.TrimSpace(val)))
	}

	return result
}
