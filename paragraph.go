package main

import(
	"fmt"
	"strings"
	"math/rand"
)

var paragraph string
fmt.Println("Enter a paragraph")
fmt.Scan(paragraph)
lowercase = strings.ToLower(paragraph)
splitStrings = strings.Split(lowercase, " ") //String array
fmt.Println("Number of words: ", len(splitStrings))
var distinct[] string
included := false
for i := 0; i < len(splitStrings); i++ {
	for j := 0; j < len(distinct); j++ {
		if splitStrings[i] == splitStrings[j] {
			included = true
		}
	}
	if (!included) {
		ndistinct = append(distinct, splitStrings[i])
	}
}
fmt.Println("Number of distict words: ", len(distict))

func main() {
	
}