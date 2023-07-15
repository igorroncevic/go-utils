// Inspired by https://github.com/zyedidia/generic and his implementation of the hashmap.
// Though, some of it had to be adapted for my own better understanding.
package hashmap

import "github.com/igorroncevic/go-utils/util"

type Map[K, V any] struct {
	entries  []entry[K, V]
	capacity uint64
	length   uint64
	readonly bool

	ops ops[K]
}

type entry[K, V any] struct {
	key    K
	filled bool
	value  V
}

type ops[T any] struct {
	equals func(a, b T) bool
	hash   func(t T) uint64
}

// New constructs a new map with the given capacity.
func New[K, V any](capacity uint64, equals util.EqualsFn[K], hash util.HashFn[K]) *Map[K, V] {
	if capacity == 0 {
		capacity = 1
	}
	capacity = pow2ceil(capacity)

	return &Map[K, V]{
		entries:  make([]entry[K, V], capacity),
		capacity: capacity,
		ops: ops[K]{
			equals: equals,
			hash:   hash,
		},
	}
}

// Get returns the value stored for this key, or false if there is no such
// value.
func (m *Map[K, V]) Get(key K) (V, bool) {
	idx := m.getIndex(m.ops.hash(key)) // Possible index

	// Similar hashes should be close to each other, so iterate through them
	for m.entries[idx].filled {
		if m.ops.equals(m.entries[idx].key, key) {
			return m.entries[idx].value, true
		}

		idx++
		if idx >= m.capacity { // Revert back to start
			idx = 0
		}
	}

	var empty V
	return empty, false
}

func (m *Map[K, V]) resize(newcap uint64) {
	newm := Map[K, V]{
		capacity: newcap,
		length:   m.length,
		entries:  make([]entry[K, V], newcap),
		ops:      m.ops,
	}

	for _, ent := range m.entries {
		if ent.filled {
			newm.Put(ent.key, ent.value) // redistribute old values in new map
		}
	}

	m.capacity = newm.capacity
	m.entries = newm.entries
}

// Put maps the given key to the given value. If the key already exists its
// value will be overwritten with the new value.
func (m *Map[K, V]) Put(key K, val V) {
	if m.length >= m.capacity/2 {
		m.resize(m.capacity * 2)
	} else if m.readonly {
		entries := make([]entry[K, V], len(m.entries), cap(m.entries))
		copy(entries, m.entries)
		m.entries = entries
		m.readonly = false
	}

	idx := m.getIndex(m.ops.hash(key)) // Possible index

	// Iterate until current element is no longer filled
	for m.entries[idx].filled {
		if m.ops.equals(m.entries[idx].key, key) { // found this exact key, just update the value
			m.entries[idx].value = val
			return
		}

		idx++
		if idx >= m.capacity {
			idx = 0
		}
	}

	// Put the value at this index, since it's empty
	m.entries[idx].key = key
	m.entries[idx].value = val
	m.entries[idx].filled = true
	m.length++
}

func (m *Map[K, V]) remove(idx uint64) {
	var k K
	var v V
	m.entries[idx].filled = false
	m.entries[idx].key = k
	m.entries[idx].value = v
	m.length--
}

// Remove removes the specified key-value pair from the map.
func (m *Map[K, V]) Remove(key K) {
	idx := m.getIndex(m.ops.hash(key)) // Possible index

	// Iterate through elements that aren't the one we are looking for
	for m.entries[idx].filled && !m.ops.equals(m.entries[idx].key, key) {
		idx = (idx + 1) & (m.capacity - 1)
	}

	// Element not found
	if !m.entries[idx].filled {
		return
	}

	if m.readonly {
		entries := make([]entry[K, V], len(m.entries), cap(m.entries))
		copy(entries, m.entries)
		m.entries = entries
		m.readonly = false
	}

	m.remove(idx)

	idx = (idx + 1) & (m.capacity - 1)
	for m.entries[idx].filled { // redistribute old nearby values, since they now can be at more adequate spots
		keyRehash := m.entries[idx].key
		valRehash := m.entries[idx].value

		m.remove(idx)
		m.Put(keyRehash, valRehash)

		idx = (idx + 1) & (m.capacity - 1)
	}

	// halves the array if it is 1/8 full or less
	if m.length > 0 && m.length <= m.capacity/8 {
		m.resize(m.capacity / 2)
	}
}

// Clear removes all key-value pairs from the map.
func (m *Map[K, V]) Clear() {
	for idx, entry := range m.entries {
		if entry.filled {
			m.remove(uint64(idx))
		}
	}
}

// Size returns the number of items in the map.
func (m *Map[K, V]) Size() int {
	return int(m.length)
}

// Copy returns a copy of this map. The copy will not allocate any memory until
// the first write, so any number of read-only copies can be made without any
// additional allocations.
func (m *Map[K, V]) Copy() *Map[K, V] {
	m.readonly = true
	return &Map[K, V]{
		entries:  m.entries,
		capacity: m.capacity,
		length:   m.length,
		readonly: true,
		ops:      m.ops,
	}
}

// Each calls 'fn' on every key-value pair in the hashmap in no particular
// order.
func (m *Map[K, V]) Each(fn func(key K, val V)) {
	for _, ent := range m.entries {
		if ent.filled {
			fn(ent.key, ent.value)
		}
	}
}

// getIndex calculates possible index based on hash and hashmap capacity.
func (m *Map[K, V]) getIndex(hash uint64) uint64 {
	return hash & (m.capacity - 1)
}

// pow2ceil helps determine capacity, which is always to the power of 2.
func pow2ceil(num uint64) uint64 {
	power := uint64(1)
	for power < num {
		power *= 2
	}
	return power
}
