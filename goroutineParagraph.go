package main

import(
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func reverse(words[] string) {
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i], " ")
	}
	fmt.Println("")
}

func count(words[] string) {
	fmt.Println("Word count: ", len(words))
}

func main() {
	var paragraph string
	fmt.Scan(paragraph)
	words := strings.Split(paragraph, " ")
	go reverse(words)
	go count(words)
	time.Sleep(time.Second)
}