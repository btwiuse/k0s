package manager

import (
	"github.com/btwiuse/conntroll/pkg"
	"github.com/btwiuse/gods/maps/linkedhashmap"
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
}

func (m *manager) Del(id string) {
	m.Remove(id)
}

func (m *manager) Has(id string) bool {
	_, ok := m.Map.Get(id)
	return ok
}

func (m *manager) Get(id string) pkg.Tider {
	v, _ := m.Map.Get(id)
	return v.(pkg.Tider)
}

func (m *manager) Add(idr pkg.Tider) {
	m.Map.Put(idr.ID(), idr)
}

func (m *manager) Values() (ret []pkg.Tider) {
	vals := m.Map.Values()
	for _, v := range vals {
		ret = append(ret, v.(pkg.Tider))
	}
	return ret
}
