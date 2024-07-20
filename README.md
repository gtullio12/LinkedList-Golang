# Linked List in Golang

Simple linked list built using Golang. The linked list type has 4 methods; Add(n Node), RemoveTail(), RemoveHead(), and PrintLinkedList().

An example might be: \
```
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
        linkedList.Add(*nodeToAdd) // (123)->(999)
        linkedList.RemoveTail() // (123)
        linkedList.RemoveHead() // ()

```
