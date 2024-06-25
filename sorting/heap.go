package main

import "fmt"

type MaxHeap struct {
	slice    []int
	heapSize int
}

func (h *MaxHeap) heapify(length, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && h.slice[left] > h.slice[largest] {
		largest = left
	}

	if right < length && h.slice[right] > h.slice[largest] {
		largest = right
	}

	if largest != i {
		h.slice[i], h.slice[largest] = h.slice[largest], h.slice[i]
		h.heapify(length, largest)
	}
}

func (h *MaxHeap) sort() {
	length := len(h.slice)

	for i := length/2 - 1; i >= 0; i-- {
		h.heapify(length, i)
	}

	for i := length - 1; i >= 0; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapify(i, 0)
	}
}

// main function to test the MaxHeap sorting
func main() {
	arr := []int{5, 3, 2, 1, 0, 4}
	h := MaxHeap{slice: arr}

	fmt.Println("Original array:", h.slice)
	h.sort()
	fmt.Println("Sorted array:", h.slice)
}
