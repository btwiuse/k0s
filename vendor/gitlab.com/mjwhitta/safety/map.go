package safety

import "sync"

// Map is a thread-safe map[interface{}]interface{} implementation.
type Map struct {
	sync.RWMutex
	themap map[interface{}]interface{}
}

// NewMap will return a pointer to a new Map instance.
func NewMap() *Map {
	return &Map{
		themap: map[interface{}]interface{}{},
	}
}

// Clear will delete all map entries
func (m *Map) Clear() {
	m.Lock()
	defer m.Unlock()

	for k := range m.themap {
		delete(m.themap, k)
	}
}

// Delete will delete a map entry and return the deleted entry.
func (m *Map) Delete(k interface{}) (v interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.themap[k]; ok {
		v = m.themap[k]
		delete(m.themap, k)
	}

	return
}

// Get will return a map entry.
func (m *Map) Get(k interface{}) (v interface{}, ok bool) {
	m.RLock()
	defer m.RUnlock()

	if _, ok = m.themap[k]; ok {
		v = m.themap[k]
	}

	return
}

// Keys will return a snapshot of the valid keys. There is no
// guarantee that the keys will remain valid for any amount of time.
func (m *Map) Keys() (keys []interface{}) {
	m.RLock()
	defer m.RUnlock()

	for k := range m.themap {
		keys = append(keys, k)
	}

	return
}

// Put will put a new entry into the map.
func (m *Map) Put(k interface{}, v interface{}) {
	m.Lock()
	defer m.Unlock()

	m.themap[k] = v
}

// Range will loop over the map and run the specified function for
// each entry. The return value determines whether or not to break the
// loop. You should not add or delete entries within Range, and you
// should avoid calling other Map functions or you may cause deadlock.
// Range should be safe to nest for any read operations.
func (m *Map) Range(f func(key, val interface{}) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.themap {
		if f(k, v) {
			break
		}
	}
}

// RangeChange will loop over the map and run the specified function
// for each entry, storing the first return value as the new entry
// value. The second return value determines whether or not to break
// the loop. You should not add or delete entries within RangeChange,
// and you should avoid calling other Map functions or you may cause
// deadlock.
func (m *Map) RangeChange(
	f func(key, val interface{}) (interface{}, bool),
) {
	var stop bool

	m.Lock()
	defer m.Unlock()

	for k, v := range m.themap {
		m.themap[k], stop = f(k, v)

		if stop {
			break
		}
	}
}
