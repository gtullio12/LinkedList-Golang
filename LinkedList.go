package main

import "fmt"

type Node struct {
	value any
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Add(n Node) {
	if ll.head == nil {
		ll.head = &n
	} else {
		runner := ll.head
		// Go to last element in linked list
		for runner.next != nil {
			runner = runner.next
		}
		runner.next = &n
	}
}

func (ll *LinkedList) RemoveTail() {
	if ll.head == nil {
		return
	} else if ll.head.next == nil {
		ll.head = nil
	} else {
		runner := ll.head

		for runner.next.next != nil {
			runner = runner.next
		}
		runner.next = nil
	}
}

func (ll *LinkedList) RemoveHead() {
	if ll.head == nil {
		return
	} else if ll.head.next == nil {
		ll.head = nil
	} else {
		ll.head = ll.head.next
	}
}

func (ll LinkedList) PrintLinkedList() {
	if ll.head == nil {
		fmt.Println("(/)")
	} else {
		runner := ll.head

		for runner.next != nil {
			fmt.Printf("(%v)->", runner.value)
			runner = runner.next
		}
		fmt.Printf("(%v)\n", runner.value)
	}

}
