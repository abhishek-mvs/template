package db


import (
	"sync"
)

type Redis struct {
	data     map[string]interface{}
	mu       sync.RWMutex
}

func NewRedis() *Redis {
	return &Redis{
		data:     make(map[string]interface{}),
	}
}

// SET key value
func (r *Redis) Set(key string, value interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[key] = value
}

// GET key
func (r *Redis) Get(key string) (interface{}, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, ok := r.data[key]
	return val, ok
}

// DEL key
func (r *Redis) Del(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.data, key)
}
