package main

import (
	"fmt"
	"strings"
)

type Node struct {
	prev *Node
	data int
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (d DoublyLinkedList) String() string {
	sb := strings.Builder{}
	for iterator := d.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%d", iterator.data))
		if iterator.next != nil {
			sb.WriteString(" -> ")
		}
	}
	return sb.String()
}

func (d *DoublyLinkedList) Append(value int) {
	node := &Node{data: value}
	if d.head == nil {
		d.head = node
		d.tail = node
		return
	}

	node.prev = d.tail
	d.tail.next = node
	d.tail = node
}

func main() {
	dll := &DoublyLinkedList{}
	dll.Append(4)
	dll.Append(2)
	dll.Append(7)
	dll.Append(5)
	dll.Append(6)

	fmt.Println(dll)
}
