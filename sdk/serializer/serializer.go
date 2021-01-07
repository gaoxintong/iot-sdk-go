package serializer

// Serializer 序列化
type Serializer interface {
	Marshal(data []interface{}, t Type) ([]byte, error)
	Unmarshal(data []byte, t Type) (interface{}, error)
}

// Type 数据交互类型
type Type int

// 数据交互类型
const (
	PostProperty = iota
	OnProperty
	PostEvent
	OnCommand
)
