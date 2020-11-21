package express

import "reflect"

type Result []interface{}

func newResult() Result {
	return make(Result, 0)
}

func (this Result) IsEmpty() bool {
	return len(this) == 0
}

func (this Result) Len() int {
	return len(this)
}

func result(values []reflect.Value) Result {
	ret := newResult()
	if values == nil || len(values) == 0 {
		return ret
	}
	for _, v := range values {
		ret = append(ret, v.Interface())
	}
	return ret
}