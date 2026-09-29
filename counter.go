// Package counter implements a small in-memory hit counter store.
package counter

import "sync"

// Store holds named hit counts, safe for concurrent single-key access.
type Store struct {
	mu     sync.Mutex
	counts map[string]int
}

// NewStore creates an empty Store.
func NewStore() *Store {
	return &Store{counts: make(map[string]int)}
}

// Increment adds 1 to key's count and returns the new value.
func (s *Store) Increment(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counts[key]++
	return s.counts[key]
}

// Get returns key's current count.
func (s *Store) Get(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.counts[key]
}

// TODO: IncrementBatch is not implemented yet. See task brief.
func (s *Store) IncrementBatch(keys []string) {
	panic("IncrementBatch not implemented")
}
