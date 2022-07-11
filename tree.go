package main

import(
	"fmt"
	"strings"
	"math/rand"
)

type Node struct {
	value int
	left *Node
	right *Node
}

type Tree struct {
	root *Node
	size int
	depth int
}

var tree Tree

func iterate(value int, node *Node) *Node {
	if value < node.value {
		if node.left == null {
			return node
		} else {
			iterate(value, node.left)
		}
	} else if value > node.value {
		if node.right == null {
			return node
		} else {
			iterate(value, node.right)
		}
	}
}

func add(value int) {
	newNode := Node{value, null, null}
	if tree == null {
		tree = Tree{null, 0, 0}
		tree.root = newNode
		tree.size++
		tree.depth++
	} else {
		node = iterate(value, tree.root)
		if value < node.value {
			node.left = newNode
		} else if value > node.value {
			node.right = newNode
		}
		tree.size++
	}
}

func makeSlice(node *Node, slice []int) {
	if node.left != null {
		makeArray(node.left)
	}

	slice = append(slice, node.value)

	if node.right != null {
		makeArray(node.right)
	}
}

func isBinaryTree() bool {
	var slice []int
	makeArray(tree.root, slice)
	cur = slice[0]
	for i := 1; i < len(size); i++ {
		if slice[i] < cur {
			return false
		} else {
			cur = slice[i]
		}
	}
	return true
}

func main() {

}