package main

import (
	"fmt"

	"k0s.io/pkg/rng"
)

func main() {
	fmt.Println(rng.New())
}
