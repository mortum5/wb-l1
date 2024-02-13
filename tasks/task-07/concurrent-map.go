package main

import (
	"fmt"
	"sync"
)

/**
 *Реализовать конкурентную запись данных в map
 */

const (
	N int = 5
)

var (
	wg sync.WaitGroup
)

type SafeMap[K comparable, V any] struct {
	mu   *sync.RWMutex
	data map[K]V
}

// New
func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		mu:   new(sync.RWMutex),
		data: make(map[K]V),
	}
}

// Insert
func (m *SafeMap[K, V]) Insert(k K, v V) {
	m.mu.Lock()
	m.data[k] = v
	m.mu.Unlock()
}

// Get
func (m *SafeMap[K, V]) Get(k K) (V, error) {
	m.mu.RLock()
	val, ok := m.data[k]
	m.mu.RUnlock()

	if !ok {
		return val, fmt.Errorf("key %v not found", k)
	}
	return val, nil
}

// Print
func (m *SafeMap[K, V]) Print() {
	m.mu.RLock()
	fmt.Println(m.data)
	m.mu.RUnlock()
}

// Length
func (m *SafeMap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

// HasKey
func (m *SafeMap[K, V]) HasKey(k K) bool {
	m.mu.RLock()
	_, ok := m.data[k]
	m.mu.RUnlock()
	return ok
}

func (m *SafeMap[K, V]) Delete(k K, v V) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[k]; !ok {
		return fmt.Errorf("key %v not found", k)
	}

	delete(m.data, k)

	return nil
}

func main() {
	sm := New[int, int]()
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < 10; j++ {
				sm.Insert(i, j%(i+1))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
