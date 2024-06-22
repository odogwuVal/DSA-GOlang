package main

import (
	"fmt"
	"strings"
)

type node struct {
	number int
	next   *node
}

type linkedList struct {
	head *node
	len  int
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.number)
}

func (l *linkedList) insert(num int) {
	newNode := new(node)
	newNode.number = num
	newNode.next = nil

	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for ; iterator.next != nil; iterator = iterator.next {

		}
		iterator.next = newNode
	}

}

func (l linkedList) Get(num int) *node {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.number == num {
			return iterator
		}
	}
	return nil
}

func (l *linkedList) Remove(num int) {
	if l.head.number == num {
		l.head = l.head.next
	} else {
		var previous *node
		for current := l.head; current != nil; current = current.next {
			if current.number == num {
				previous.next = current.next
				return
			}
			previous = current
		}
	}
}

func (l *linkedList) RemovePosition(position int) {
	if position == 0 {
		l.head = l.head.next
	} else {
		iterator := l.head
		for i := 0; i < position-1; i++ {
			iterator = iterator.next
		}
		iterator.next = iterator.next.next
	}
}

func (l *linkedList) Prepend(num int) {
	newNode := &node{number: num}
	temp := l.head
	l.head = newNode
	l.head.next = temp
}

func (l *linkedList) InsertPosition(num int, position int) {
	newNode := &node{number: num}
	if position == 0 {
		l.Prepend(num)
	} else {
		iterator := l.head
		for i := 0; i < position-1; i++ {
			iterator = iterator.next
		}
		newNode.next = iterator.next
		iterator.next = newNode
	}
}

func (l linkedList) String() string {
	sb := strings.Builder{}
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%s", iterator))
		if iterator.next != nil {
			sb.WriteString(" -> ")
		}
	}
	return sb.String()
}

func main() {
	l := linkedList{}
	l.insert(1)
	l.insert(2)
	l.insert(4)
	l.insert(5)
	// l.Prepend(3)
	// l.InsertPosition(7, 3)
	// l.Remove(1)
	l.RemovePosition(3)
	fmt.Println(l)
	// fmt.Println(l.Get(1))
}
