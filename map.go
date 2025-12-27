package goutils

import "sync"

// SafeMap은 멀티 쓰레드 환경에서 안전하게 사용할 수 있는 맵입니다.
type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// ReverseMap는 주어진 맵의 키와 값을 뒤바꾼 새로운 맵을 반환합니다.
func ReverseMap[T, U comparable](m map[T]U) map[U]T {
	res := make(map[U]T)
	for k, v := range m {
		res[v] = k
	}
	return res
}

// NewSafeMap는 새로운 SafeMap 인스턴스를 생성합니다.
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Get은 주어진 키에 해당하는 값을 반환합니다.
func (sm *SafeMap[K, V]) Get(key K) (V, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}

// Set은 주어진 키에 값을 설정합니다.
func (sm *SafeMap[K, V]) Set(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Delete는 주어진 키에 해당하는 항목을 삭제합니다.
func (sm *SafeMap[K, V]) Delete(key K) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SafeMap[K, V]) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

func (sm *SafeMap[K, V]) Keys() []K {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	keys := make([]K, 0, len(sm.data))
	for k := range sm.data {
		keys = append(keys, k)
	}
	return keys
}

func (sm *SafeMap[K, V]) Values() []V {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	values := make([]V, 0, len(sm.data))
	for _, v := range sm.data {
		values = append(values, v)
	}
	return values
}
