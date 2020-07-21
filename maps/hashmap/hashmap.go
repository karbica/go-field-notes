package hashmap

// Map holds the state of the map instance.
type Map struct {
	entries map[interface{}]interface{}
}

// New returns a new instance of a hash map.
func New() *Map {
	entries := make(map[interface{}]interface{})
	return &Map{entries}
}

// Size returns the size of the map.
func (m *Map) Size() (size int) {
	return len(m.entries)
}

// Empty returns whether or not the hash map is empty.
func (m *Map) Empty() (ok bool) {
	return m.Size() == 0
}

// Get retrieves an entry from the hash map.
func (m *Map) Get(key interface{}) (value interface{}, ok bool) {
	value, ok = m.entries[key]
	return value, ok
}

// Set add an entry to the hash map.
func (m *Map) Set(key interface{}, value interface{}) {
	m.entries[key] = value
}

// Delete removes an entry from the hash map.
func (m *Map) Delete(key interface{}) {
	delete(m.entries, key)
}

// Keys returns a slice of keys in the hash map.
// The keys returned are unordered.
func (m *Map) Keys() (keys []interface{}) {
	keys = make([]interface{}, m.Size())
	i := 0

	for key := range m.entries {
		keys[i] = key
		i++
	}

	return keys
}
