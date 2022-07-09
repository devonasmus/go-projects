package main

import(
	"fmt"
	"strings"
	"math/rand"
)

func highest(numbers ... int) int {
	highest := 0
	for cur := range numbers {
		if cur > highest {
			highest = cur
		}
	}
	return highest
}

func main() {
	
}