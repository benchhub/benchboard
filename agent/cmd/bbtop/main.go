package main

import (
	"fmt"
	"github.com/benchhub/benchboard/agent/pkg/host"
)

func main() {
	fmt.Println("something works like top")
	c := host.CPUStat{}
	fmt.Println(c.Guest)
}
