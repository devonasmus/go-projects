package main

import(
	"fmt"
	"math/rand"
	"time"
)

func average(slice[] int, startIndex int, endIndex int) {
	average := 0
	for i := startIndex; i <= endIndex; i++ {
		average += slice[i]
	}
	average /= ((endIndex - startIndex) + 1)
	fmt.Println("Average: ", average)
}

func main() {
	var slice[] int
	for i := 0; i < 20; i++ {
		slice = append(slice, rand.Intn(100))
	}
	go average(slice, 0, 4)
	go average(slice, 5, 9)
	go average(slice, 10, 14)
	go average(slice, 15, 19)

	time.Sleep(time.Second)
}