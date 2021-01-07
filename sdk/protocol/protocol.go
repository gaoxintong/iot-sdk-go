package protocol

// Protocol 协议
type Protocol interface {
	Publish(params map[string]interface{}) error
	Subscribe(params map[string]interface{}) error
	Unsubscribe(params map[string]interface{}) error
	MakeOpts(params map[string]interface{}) interface{}
	NewClient(opts interface{}) error
	GetName() string
	GetInstance() interface{}
}
