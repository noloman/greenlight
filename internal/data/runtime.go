package data

import (
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = fmt.Errorf("invalid runtime format")

// MarshalJSON is a custom JSON marshaler for the Runtime type.
// This method is called by the json.Marshal() function and it returns the JSON representation of the Runtime type.
// It needs to be wrapped in double quotes, otherwise it won't be interpreted as a JSON String and
// we'll receive a runtime error similar to:\n
// json: error calling MarshalJSON for type data.Runtime: invalid character 'm' after top-level value
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

// Implement a UnmarshalJSON() method on the Runtime type so that it satisfies the
// json.Unmarshaler interface. IMPORTANT: Because UnmarshalJSON() needs to modify the
// receiver (our Runtime type), we must use a pointer receiver for this to work
// correctly. Otherwise, we will only be modifying a copy (which is then discarded when
// this method returns).
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.Atoi(parts[0])
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)
	return nil
}
