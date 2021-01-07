package request

// Request 请求
type Request struct {
	Topic    string
	Qos      byte
	Retained bool
	Payload  interface{}
	Callback func(err error)
}
