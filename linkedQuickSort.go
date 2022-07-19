package main

import(
	"fmt"
	"math/rand"
)

type Node struct {
	value int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func sort(list *List) {
	if list.head != nil && list.head.next != nil {
		pivot := list.head
		leftList := List{nil, nil}
		rightList := List{nil, nil}
		
		for list.head.next != nil {
			currentNode := list.head.next
			list.head.next = list.head.next.next
			currentNode.next = nil
			if currentNode.value < pivot.value {
				if leftList.head == nil {
					leftList.head = currentNode
					leftList.tail = currentNode
				} else {
					leftList.tail.next = currentNode
					leftList.tail = currentNode
				}
			} else {
				if rightList.head == nil {
					rightList.head = currentNode
					rightList.tail = currentNode
				} else {
					rightList.tail.next = currentNode
					rightList.tail = currentNode
				}
			}
		}

		sort(&leftList)
		sort(&rightList)

		pivot.next = rightList.head

		if leftList.head != nil {
			leftList.tail.next = pivot
			list.head = leftList.head
		} else {
			list.head = pivot
		}

		if rightList.tail != nil {
			list.tail = rightList.tail
		} else {
			list.tail = pivot
		}
	}
}

func main() {
	node := Node{rand.Intn(250), nil}
	list := List{&node, &node}
	currentNode := list.head
	for i := 1; i < 250; i++ {
		value := rand.Intn(250)
		newNode := Node{value, nil}
		currentNode.next = &newNode
		currentNode = currentNode.next
		list.tail = &newNode
	}
	fmt.Println("Test")
	sort(&list)

	currentNodetoPrint := list.head
	for currentNodetoPrint.next != nil {
		fmt.Println(currentNodetoPrint.value)

		currentNodetoPrint = currentNodetoPrint.next
	}
	fmt.Println(currentNodetoPrint.value)
}