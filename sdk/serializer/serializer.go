package serializer

// Serializer 序列化
type Serializer interface {
	Marshal(data interface{}) (interface{}, error)
	Unmarshal(data interface{}) (interface{}, error)
	MakePropertyData(data *Property) ([]byte, error)
	MakeEventData(data *Property) ([]byte, error)
	UnmarshalCommand(data []byte) (*Command, error)
}

// Property 属性
type Property struct {
	SubDeviceID uint16
	PropertyID  uint16
	Value       []interface{}
}

// Command 命令
type Command struct {
	ID          uint16
	SubDeviceID uint16
	Params      map[int]interface{}
}
