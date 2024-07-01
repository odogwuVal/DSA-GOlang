package main

import "fmt"

// create a count array of size max+1, where max is the maximum element in the input array
// returns the maximum value of the array
func max(arr []int) int {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

// count the occurrences of each element in the input array and store it in the count array.
func countOccurrence(arr []int, count []int) {
	for _, v := range arr {
		count[v]++
	}
}

// update the count array by adding the previous counts.
func updateCount(count []int) {
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
}

// Finally, we build the sorted array using the count array and the input array.
func buildSorted(arr []int, count []int) []int {
	sorted := make([]int, len(arr))
	for _, v := range arr {
		count[v]--
		sorted[count[v]] = v
	}
	return sorted
}

func CountingSort(arr []int) []int {
	max := max(arr)
	count := make([]int, max+1)
	countOccurrence(arr, count)
	updateCount(count)
	return buildSorted(arr, count)
}

func main() {
	arr := []int{4, 2, 2, 8, 3, 3, 1}
	fmt.Println("Original array:", arr)
	sortedArr := CountingSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
