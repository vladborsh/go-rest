package storage

import (
	"errors"
	"sync"
)

type DB interface {
	Get(key string) ([]byte, error)
	Put(key string, val []byte) error
}

type Storage struct {
	m     map[string][]byte
	mutex sync.RWMutex
}

func NewDataBase() DB {
	return &Storage{m: make(map[string][]byte)}
}

func (s *Storage) Get(key string) ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	v, ok := s.m[key]
	if !ok {
		return nil, errors.New("Not found")
	}
	return v, nil
}

func (s *Storage) Put(key string, val []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m[key] = val
	return nil
}
