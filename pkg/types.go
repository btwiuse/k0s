package pkg

import (
	"time"
)

type Timer interface {
	Time() time.Time
}

type Namer interface {
	Name() string
}

type IDer interface {
	ID() string
}

type Tider interface {
	Namer
	Timer
	IDer
}

type Manager interface {
	Add(Tider)
	Del(string)
	Has(string) bool
	Get(string) Tider
	Values() []Tider
	Size() int
}
