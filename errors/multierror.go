package errors

import (
	"bytes"
	"fmt"
)

type ErrorString string

func (es ErrorString) Error() string {
	return string(es)
}

type MultiError []error

func (me *MultiError) Append(err interface{}) (added bool) {
	tmp := *me
	added = true
	switch err := err.(type) {
	case error:
		tmp = append(tmp, err)
	case string:
		tmp = append(tmp, ErrorString(err))
	case []byte:
		tmp = append(tmp, ErrorString(err))
	case *MultiError:
		tmp = append(tmp, (*err)...)
	case MultiError:
		tmp = append(tmp, err...)
	case []error:
		tmp = append(tmp, err...)
	case nil:
		added = false
	default:
		added = false
	}
	*me = tmp
	return
}

func (me *MultiError) Len() int {
	return len(*me)
}

func (me *MultiError) Error() string {
	tmp := *me
	if len(tmp) == 1 {
		return tmp[0].Error()
	}

	buf := &bytes.Buffer{}
	buf.WriteString("Multiple Errors Found:\n")
	for i, e := range tmp {
		fmt.Fprintf(buf, "\t[%d] %s\n", i, e.Error())
	}
	return buf.String()[:buf.Len()-1] // skip the last \n
}
