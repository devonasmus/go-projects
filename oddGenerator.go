package main

import(
	"fmt"
	"strings"
	"math/rand"
)

func oddGenerator() func() int {
	i := 1
	return func() int {
		i += 2
		return i
	}
}

func main() {
	
}