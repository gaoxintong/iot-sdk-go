package storage

// Storage 存储
type Storage interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Del(key string) error
}
