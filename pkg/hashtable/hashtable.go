package hashtable

type HashTable[Key comparable, Value any] struct {
	key2indexes map[Key]int
	data        []Value
	cap         int
}

func (t *HashTable[Key, Value]) initialize() {
	t.key2indexes = make(map[Key]int, t.cap)
	t.data = make([]Value, 0, t.cap)
}

func New[Key comparable, Value any](opts ...Option[Key, Value]) *HashTable[Key, Value] {
	table := &HashTable[Key, Value]{}
	for _, opt := range opts {
		opt(table)

	}
	table.initialize()
	return table
}

type Option[Key comparable, Value any] func(*HashTable[Key, Value])

func WithCapacity[K comparable, V any](cap int) Option[K, V] {
	return func(t *HashTable[K, V]) {
		t.cap = cap
	}
}
