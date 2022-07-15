package main

import(
	"fmt"
)

type Hero struct {
	health int
	attack int
	defense int
	potion [5]int
	gold int
}

type Goblin struct {
	health int
	attack int
	defense int
}