package bahn

type CacheBackend interface {
	Set(key string, value interface{}) error
	Get(key string, value interface{}) error
}
