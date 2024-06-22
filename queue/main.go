package main

import "fmt"

type Queue []string

func (q Queue) isEmpty() bool {
	return len(q) == 0
}

func (q *Queue) Enqueue(str string) {
	*q = append(*q, str)
}

func (q *Queue) Dequeue() (string, bool) {
	if q.isEmpty() {
		return "", false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func main() {
	var queue Queue
	queue.Enqueue("Everyboody")
	queue.Enqueue("Loves")
	queue.Enqueue("Val")
	for !queue.isEmpty() {
		val, ok := queue.Dequeue()
		if ok == true {
			fmt.Println(val)
		}
	}
}
