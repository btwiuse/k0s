package rng

import (
	"math/rand"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

func New() string {
	rand.Seed(time.Now().UnixNano())
	return namesgenerator.GetRandomName(0)
}
