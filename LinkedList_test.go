package main

import (
	"testing"
)

func TestLinkedListCreation(t *testing.T) {
	var node3 *Node = &Node{
		value: 789,
		next:  nil,
	}
	var node2 *Node = &Node{
		value: 345,
		next:  node3,
	}
	var node1 *Node = &Node{
		value: 123,
		next:  node2,
	}
	var linkedList *LinkedList = &LinkedList{
		head: node1,
	}

	if linkedList.head.value != 123 {
		t.Errorf("Linked list head should be %v, but was %v\n", 123, linkedList.head.value)
	}
	if linkedList.head.next.value != 345 {
		t.Errorf("Second node should have value: %v, but was %v\n", 345, linkedList.head.next.value)
	}
}

func TestLinkedListAdd(t *testing.T) {
	var head *Node = &Node{
		value: 123,
		next:  nil,
	}
	var nodeToAdd *Node = &Node{
		value: 999,
		next:  nil,
	}
	var linkedList *LinkedList = &LinkedList{
		head: head,
	}
	linkedList.Add(*nodeToAdd)

	if linkedList.head.next == nil {
		t.Errorf("Failed to add node")
	}

	if linkedList.head.next.value != 999 {
		t.Errorf("Second node should have value: %v, but was %v\n", 999, linkedList.head.next.value)
	}

}

func TestEmptyLinkedListAdd(t *testing.T) {
	var nodeToAdd *Node = &Node{
		value: 123,
		next:  nil,
	}
	var linkedList *LinkedList = &LinkedList{
		head: nil,
	}

	linkedList.Add(*nodeToAdd)

	linkedList.Add(*&Node{
		value: 6777,
		next:  nil,
	})

	if linkedList.head == nil {
		t.Error("Failed to add node to empty linked list")
	}

	if linkedList.head.value != 123 {
		t.Errorf("Head node should have value %v, but was %v\n", 123, linkedList.head.value)
	}
}

func TestLinkedListRemoveTail(t *testing.T) {
	var node3 *Node = &Node{
		value: 789,
		next:  nil,
	}
	var node2 *Node = &Node{
		value: 345,
		next:  node3,
	}
	var node1 *Node = &Node{
		value: 123,
		next:  node2,
	}
	var linkedList *LinkedList = &LinkedList{
		head: node1,
	}

	linkedList.RemoveTail()

	if linkedList.head.next.next != nil {
		t.Error("Failed to remove tail")
	}
}

func TestLinkedListWithOnlyHeadRemoveTail(t *testing.T) {
	var node *Node = &Node{
		value: 123,
		next:  nil,
	}
	var linkedList *LinkedList = &LinkedList{
		head: node,
	}

	linkedList.RemoveTail()

	if linkedList.head != nil {
		t.Error("Failed to remove head")
	}
}

func TestLinkedListRemoveHead(t *testing.T) {
	var node3 *Node = &Node{
		value: 789,
		next:  nil,
	}
	var node2 *Node = &Node{
		value: 345,
		next:  node3,
	}
	var node1 *Node = &Node{
		value: 123,
		next:  node2,
	}
	var linkedList *LinkedList = &LinkedList{
		head: node1,
	}

	linkedList.RemoveHead()

	if linkedList.head.value != 345 {
		t.Errorf("Failed to remove head, head should be %v, but was %v\n", 345, linkedList.head.value)
	}
}

func TestLinkedListSizeOneRemoveHead(t *testing.T) {
	var node *Node = &Node{
		value: 123,
		next:  nil,
	}
	var linkedList *LinkedList = &LinkedList{
		head: node,
	}

	linkedList.RemoveHead()

	if linkedList.head != nil {
		t.Error("Failed to remove head")
	}
}
