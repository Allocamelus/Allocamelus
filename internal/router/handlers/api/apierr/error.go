package apierr

import (
	"errors"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// APIErr struct with error
type APIErr struct {
	Err interface{} `json:"error"`
}

// New Create An json compatable response error
func New(data interface{}) APIErr {
	return APIErr{Err: data}
}

// AppendStr append string to Err if is string
func (e *APIErr) AppendStr(s string) error {
	switch e.Err.(type) {
	case string:
		var sb strings.Builder
		sb.WriteString(e.Err.(string))
		sb.WriteString(s)
		e.Err = sb.String()
	case []string:
		e.Err = append(e.Err.([]string), s)
	default:
		return errors.New("apierr/error: Error appending string")
	}
	return nil
}

// MarshalStr APIErr
func (e *APIErr) MarshalStr() (string, error) {
	return json.MarshalToString(e)
}
