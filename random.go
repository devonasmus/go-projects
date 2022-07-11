package main

import(
	"fmt"
	"strings"
	"math/rand"
)

var number int
fmt.Scan(&number)
for i := 0; i < number; i++ {
	fmt.Println(rand.Intn(200)-100)
}

func main() {

}