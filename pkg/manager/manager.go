package manager

import (
	"github.com/btwiuse/conntroll/pkg/hub"

	"github.com/btwiuse/gods/maps/linkedhashmap"
)

var (
	_ hub.Manager = (*manager)(nil)
)

func NewManager() hub.Manager {
	return &manager{
		Map: linkedhashmap.New(),
	}
}

// key: string, value: IDer
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

func (m *manager) Get(id string) hub.IDer {
	v, _ := m.Map.Get(id)
	return v.(hub.IDer)
}

func (m *manager) Add(ider hub.IDer) {
	m.Map.Put(ider.ID(), ider)
}

func (m *manager) Values() (ret []hub.IDer) {
	vals := m.Map.Values()
	// TODO: document this!!
	// ret := []hub.IDer{}
	// ret := make([]hub.IDer, len(vals))
	for _, v := range vals {
		ret = append(ret, v.(hub.IDer))
	}
	return ret
}
