package store

import "fmt"

type KVStore struct {
	Data map[string]string
}

func InitializeStore() *KVStore {
	return &KVStore{
		Data: make(map[string]string),
	}
}

func (s *KVStore) Get(key string) (string, error) {
	val, ok := s.Data[key]
	if !ok {
		return "", fmt.Errorf("key does not exist")
	}
	return val, nil
}

// updates if the key-value pair already exists
func (s *KVStore) Put(key string, value string) {
	s.Data[key] = value
}

// updates if the key-value pair already exists
func (s *KVStore) Remove(key string) error {
	_, ok := s.Data[key]
	if !ok {
		return fmt.Errorf("key does not exist")
	}

	delete(s.Data, key)
	return nil
}

func (s *KVStore) Len() int {
	return len(s.Data)
}
