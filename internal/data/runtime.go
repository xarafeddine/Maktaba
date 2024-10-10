package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {

	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)
	return nil
}

func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate a string containing the movie runtime in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)

	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to be surrounded by double quotes in order to be a valid *JSON string*.
	quotedJSONValue := strconv.Quote(jsonValue)

	// Convert the quoted string value to a byte slice and return it.
	return []byte(quotedJSONValue), nil
}
