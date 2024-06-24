package main

import "fmt"

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > temp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
	return arr
}

func main() {
	fmt.Println(insertionSort([]int{5, 3, 2, 1, 0, 4}))
}
