package amf

import (
	"io"
	"reflect"
)

type Encoder struct{}

func (e *Encoder) EncodeAmf0(w io.Writer, val interface{}) (int, error) {
	if val == nil {
		return e.EncodeAmf0Null(w, true)
	}
	v := reflect.ValueOf(val)
	if v.IsValid()
}
