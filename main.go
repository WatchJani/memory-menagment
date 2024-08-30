package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	tree := NewTree(25)

	for range 2 {
		for index := range 5 {
			number := rand.Intn(100)
			fmt.Println("num", number)
			tree.Add(index, number)
		}

		tree.Read()
		fmt.Println("=========================")
		tree.Clear()
	}

}

type Node struct {
	time  time.Time
	value int
	key   int
	link  *Node
}

type Tree struct {
	root    *Node
	block   []Node
	counter int
}

func NewTree(capacity int) *Tree {
	return &Tree{
		root:  &Node{},
		block: make([]Node, capacity),
	}
}

func (t *Tree) Add(value, key int) {
	current := t.root

	for root := t.root; root.link != nil && key > root.key; root = root.link {
		current = root
	}

	newNode := &t.block[t.counter]
	*newNode = Node{
		time:  time.Now(),
		value: value,
		key:   key,
	}
	t.counter++

	rightLink := current.link
	current.link = newNode

	if rightLink != nil {
		newNode.link = rightLink
	}
}

func (t *Tree) Read() {
	for root := t.root.link; root != nil; root = root.link {
		fmt.Println(root.key)
	}
}

func (t *Tree) Clear() {
	t.root.link = nil
	t.counter = 0
}
