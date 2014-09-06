package errors

import (
	"bytes"
	"fmt"
)

type ErrorString string

func (es ErrorString) Error() string {
	return string(es)
}

type MultiError struct {
	Errors []error
}

func (me *MultiError) Append(err interface{}) (added bool) {
	added = true
	switch err := err.(type) {
	case error:
		me.Errors = append(me.Errors, err)
	case string:
		me.Errors = append(me.Errors, ErrorString(err))
	case []byte:
		me.Errors = append(me.Errors, ErrorString(err))
	case *MultiError:
		me.Errors = append(me.Errors, err.Errors...)
	case MultiError:
		me.Errors = append(me.Errors, err.Errors...)
	case []error:
		me.Errors = append(me.Errors, err...)
	case nil:
		added = false
	default:
		added = false
	}
	return
}

func (me *MultiError) Len() int {
	return len(me.Errors)
}

func (me *MultiError) Error() string {
	if len(me.Errors) == 1 {
		return me.Errors[0].Error()
	}

	buf := &bytes.Buffer{}
	buf.WriteString("Multiple Errors Found:\n")
	for i, e := range me.Errors {
		fmt.Fprintf(buf, "\t[%d] %s\n", i, e.Error())
	}
	return buf.String()[:buf.Len()-1] // skip the last \n
}
