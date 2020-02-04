// +build uuid

package main

import (
	"fmt"
	"log"

	"github.com/denisbrodbeck/machineid"
	"k0s.io/k0s/pkg/uuid"
)

func main() {
	mid, _ := machineid.ID()
	log.Println(mid)
	{
		fmt.Println(uuid.New())
		fmt.Println(uuid.New())
		fmt.Println(uuid.New())
		fmt.Println(uuid.New())
	}
	{
		fmt.Println(uuid.NewPet("asdf"))
		fmt.Println(uuid.NewPet("asdf"))
		fmt.Println(uuid.NewPet("asdf"))
		fmt.Println(uuid.NewPet("asdf"))
	}
	{
		fmt.Println(uuid.New())
		fmt.Println(uuid.New())
		fmt.Println(uuid.New())
		fmt.Println(uuid.New())
	}
}
