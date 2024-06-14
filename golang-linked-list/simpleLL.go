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

func (l linkedList) string() string {
	sb := strings.Builder{}
	for iterator := l.head; iterator.next != nil; iterator.next {
		sb.WriteString(fmt.Sprintf("%d", iterator.number))
	}
	return sb.String()
}

func main() {
	l := linkedList{}
	l.insert(3)
}
