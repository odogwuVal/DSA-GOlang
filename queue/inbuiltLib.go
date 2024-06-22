package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()
	queue.PushFront("Everybody")
	queue.PushFront("Loves")
	queue.PushFront("Val")

	for {
		front := queue.Front()
		if front != nil {
			fmt.Println(front.Value)
			queue.Remove(front)
		} else {
			fmt.Println("Queue is now empty")
			return
		}
	}
}
