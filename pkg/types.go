package pkg

import (
	"time"
)

type Timer interface {
	Time() time.Time
}

type IDer interface {
	ID() string
}

type Tider interface {
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
