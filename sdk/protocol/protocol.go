package protocol

import "reflect"

// Protocol 协议
type Protocol interface {
	Publish(opts map[string]interface{}) error
	Subscribe(opts map[string]interface{}) error
	Unsubscribe(opts map[string]interface{}) error
	MakeOpts(opts map[string]interface{}) interface{}
	NewClient(opts interface{}) error
	GetName() string
	GetInstance() interface{}
}

// OptionsFormatter 参数格式化
func OptionsFormatter(s interface{}) map[string]interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	ret := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		ret[t.Field(i).Name] = v.Field(i).Interface()
	}
	return ret
}
