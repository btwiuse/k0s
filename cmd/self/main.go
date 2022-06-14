package main

import (
	"k0s.io/pkg/hub/self"
)

func main(){
	self.Agent("http://localhost:8000")
}
