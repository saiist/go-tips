package store

type InMemoryStore struct {
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{}
}

func (s *InMemoryStore) Get(key string) string {
	return "value"
}
