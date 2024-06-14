package main

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

func main() {
	l := linkedList{}
	l.insert(3)
}
