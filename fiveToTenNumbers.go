package main

import(
	"fmt"
	"strings"
	"math/rand"
)

var numbers[] int
var num int
fmt.Println("Enter 5 numbers")
for i := 0; i < 5; i++ {
	fmt.Scan(&num)
	numbers = append(numbers, num)
}
var affirm string
for i := 0; i < 5; i++ {
	fmt.Println("Enter y to quit")
	fmt.Scan(affirm)
	if affirm == 'y' {
		break;
	}
	fmt.Scan(&num)
	numbers = append(numbers, num)
}
product := 1
sum := 0
for i := 0; i < len(numbers); i++ {
	fmt.Println(numbers[i])
	product *= numbers[i]
	sum += numbers[i]
}
fmt.Println("Product: ", product)
fmt.Println("Sum: ", sum)

func main() {
	
}