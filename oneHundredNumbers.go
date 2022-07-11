package main

import(
	"fmt"
	"strings"
	"math/rand"
)

var numbers[] int
var number int
for i := 0; i < 100; i++ {
	fmt.Scan(&number)
	numbers = append(numbers, number)
}

max := numbers[0]
min := numbers[0]
indexOfMax := 0
indexOfMin := 0
for i := 1; i < 100; i++ {
	if numbers[i] > max {
		max = numbers[i]
		indexOfMax = i
	}
	if numbers[i] < min {
		min = numbers[i]
		indexOfMin = i
	}
}

var descending[] int
for i < 0; i < len(numbers); i++ {
	descending = append(descending, numbers[i])
}
sorted := false
for !sorted {
	sorted = true
	for i := 0; i < len(descending) - 1; i++ {
		if descending[i + 1] > descending[i] {
			sorted = false
			stored := descending[i]
			descending[i] = descending[i + 1]
			descending[i + 1] = stored
		}
	}
}

var ascending[] int
for i < 0; i < len(numbers); i++ {
	ascending = append(ascending, numbers[i])
}
sorted := false
for !sorted {
	sorted = true
	for i := 0; i < len(ascending) - 1; i++ {
		if ascending[i+ 1] < ascending[i] {
			sorted = false
			stored := ascending[i]
			ascending[i] = ascending[i + 1]
			ascending[i + 1] = stored
		}
	}
}

mean := 0
for i := 0; i < len(numbers); i++ {
	mean += numbers[i]
}
mean /= len(numbers)

median := ascending[len(ascending) / 2]

var positive[] int
for i := 0; i < len(numbers) i++ {
	if numbers[i] > 0 {
		positive = append(positive, numbers[i])
	}
}

var negative[] int
for i := 0; i < len(numbers) i++ {
	if numbers[i] < 0 {
		negative = append(negative, numbers[i])
	}
}

var arr[] int
for i := 0; i < len(numbers); i++ {
	var curr[] int
	num := numbers[i]
	curr = append(curr, numbers[i])
	for j := i; j < len(numbers); j++ {
		if numbers[j] > num {
			curr = append(curr, numbers[j])
			num = numbers[j]
		}
	}
	if len(curr) > len(arr) {
		for j := 0; j < len(curr); j++ {
			arr = append(arr, curr[j])
		}
	}
}

var noDuplicates[] int
included := false
for i := 0; i < len(numbers); i++ {
	for j := 0; j < len(noDuplicates); j++ {
		if numbers[i] == noDuplicates[j] {
			included = true
		}
	}
	if (!included) {
		noDuplicates = append(noDuplicates, numbers[i])
	}
}

func main() {
	
}