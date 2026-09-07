package storage

type Storage interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}
