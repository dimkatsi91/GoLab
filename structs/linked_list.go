package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func create(x int) *Node {
	return &Node {
		data : x,
	}
}

func (node *Node) add(x int) error {
	t := &Node {
		data : x,
	}
	if node.next == nil {
		node.next = t
	}  else {
		currentNode := node.next
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = t
	}
	return nil
}


func (node *Node) isEmpty() bool {
	if(node == nil) {
		return true
	}
	return false
}

// Display linked list's elements
//
func (node *Node) display() {
	for node!=nil {
		fmt.Printf("%d ", node.data)
		node = node.next
	}
	fmt.Println("")
}

// Counter the nodes in the linked list
//
func (node *Node) count() int {
	var counter int	// initialized with 0 | same as counter := 0
	for node!=nil {
		counter += 1
		node = node.next
	}
	return counter
}

func (node *Node) sumOfElems() int {
	var sum int // 0
	for node!=nil {
		sum += node.data
		node = node.next
	}
	return sum
}

func (node *Node) search(key int) *Node {
	for node!=nil {
		if key == node.data {
			return node
		}
		node = node.next
	}
	return &Node{-1,nil}
}

func main() {
	ll := create(10)
	ll.add(20)
	ll.add(30)
	ll.add(40)
	ll.add(50)
	ll.display()
	
	if ll.isEmpty() == true {
		fmt.Println(" Linked-List is Empty ...")
	} else {
		fmt.Println(" Linked-List is not Empty ...")
	}

	fmt.Printf("  Linked list has %d elements with Total Sum = %d \n", ll.count(), ll.sumOfElems())

	// Search for a key in the list
	//
	var key int = 33
	retLL := ll.search(key)
	if retLL.data == -1 {
		fmt.Printf("Element %d Does not exist in the list ...", key)
	} else {
		fmt.Printf("Element %d is in the linked-list - %d ", key, retLL.data)
	}

}