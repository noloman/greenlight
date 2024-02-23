package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

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
