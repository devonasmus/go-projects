package main

import(
	"fmt"
	"strings"
	"math/rand"
)

func randomNumbers() [] int {
	var numbers[] int
	for i := 0; i < 100; i++ {
		numbers = append(numbers, rand.Intn(100))
	}
	return numbers
}

func sorter(numbers[] int, bool ascending) {
	if (ascending) {
		sorted := false
		for !sorted {
			sorted = true
			for i := 0; i < len(numbers) - 1; i++ {
				if numbers[i + 1] < numbers[i] {
					sorted = false
					stored := numbers[i]
					numbers[i] = numbers[i + 1]
					numbers[i + 1] = stored
				}
			}
		}
	} else {
		sorted := false
		for !sorted {
			sorted = true
			for i := 0; i < len(numbers) - 1; i++ {
				if numbers[i + 1] > numbers[i] {
					sorted = false
					stored := numbers[i]
					numbers[i] = numbers[i + 1]
					numbers[i + 1] = stored
				}
			}
		}
	}
}

func mergeSorted(numbers1[] int, numbers2[] int) [] int {
	index1 := 0
	index2 := 0
	var mergedSlice[] int
	for index1 + index2 < len(numbers1) + len(numbers2) {
		if index2 == len(numbers2) || index1 < len(numbers1) && numbers1[index1] < numbers2[index2] {
			mergedSlice = append(mergedSlice, numbers1[index1])
			index1++
		} else {
			mergedSlice = append(mergedSlice, numbers2[index2])
			index2++
		}
	}
}

func main() {
	
}