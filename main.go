package main

import (
	"fmt"

	"github.com/cihangirozmus/go-algorithms/stacks"
)

func main() {
	mst := stacks.Constructor()
	fmt.Println(mst.GetMin())
}
