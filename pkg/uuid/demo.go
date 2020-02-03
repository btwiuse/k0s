// +build uuid

package main

import (
	"fmt"

	"k0s.io/k0s/pkg/uuid"
)

func main() {
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
