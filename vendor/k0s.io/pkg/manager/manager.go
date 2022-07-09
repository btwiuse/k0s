package manager

import (
	"sync"

	"github.com/btwiuse/gods/maps/linkedhashmap"
	"k0s.io/pkg"
)

var (
	_ pkg.Manager = NewManager()
)

func NewManager() *manager {
	return &manager{
		Map: linkedhashmap.New(),
	}
}

// key: string, value: Tider
type manager struct {
	// maps.Map
	*linkedhashmap.Map
	mu sync.RWMutex
}

func (m *manager) Del(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Remove(id)
}

func (m *manager) Has(id string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.Map.Get(id)
	return ok
}

func (m *manager) Get(id string) pkg.Tider {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, _ := m.Map.Get(id)
	return v.(pkg.Tider)
}

func (m *manager) Add(idr pkg.Tider) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Map.Put(idr.ID(), idr)
}

func (m *manager) Keys() (ret []string) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := m.Map.Keys()
	for _, v := range keys {
		ret = append(ret, v.(string))
	}
	return ret
}

func (m *manager) Values() (ret []pkg.Tider) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	vals := m.Map.Values()
	for _, v := range vals {
		ret = append(ret, v.(pkg.Tider))
	}
	return ret
}
