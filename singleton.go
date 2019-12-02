package pattern

import (
	"fmt"
	"sync"
)

type hash map[string]interface{}

type singleton struct {
	items hash
	m     sync.RWMutex
}

func (s *singleton) Set(key string, data interface{}) {
	s.m.Lock()
	defer s.m.Unlock()
	println()
	s.items[key] = data
}

func (s *singleton) Get(key string) (interface{}, error) {
	s.m.RLock()
	defer s.m.RUnlock()
	item, ok := s.items[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return item, nil
}

var (
	s    *singleton
	once sync.Once
)

func Singleton() *singleton {
	once.Do(func() {
		if s == nil {
			s = &singleton{
				items: make(hash),
			}
		}
	})

	return s
}
