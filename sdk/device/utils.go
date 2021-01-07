package device

import (
	"errors"
	"reflect"
)

// HTTPIsOK 状态码是否正常
func HTTPIsOK(resp interface{}) error {
	res := reflect.ValueOf(resp)
	if res.Kind() == reflect.Struct {
		f := res.FieldByName("Code")
		if f.IsValid() {
			if f.Interface() == 0 {
				return nil
			}
			msg := res.FieldByName("Message")
			return errors.New(msg.Interface().(string))
		}
	}
	return errors.New("response format error")
}
