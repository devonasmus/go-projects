package main

import(
	"fmt"
	"math/rand"
	"time"
)

func appendNumber(slice[] int, number int) {
	slice = append(slice, number)
}

func sortSlice(slice[] int) {
	for i := 0; i < len(slice) - 1; i++ {
		minIndex := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}
		store := slice[i]
		slice[i] = slice[minIndex]
		slice[minIndex] = store
	}
}

func main() {

}