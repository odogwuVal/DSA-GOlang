package main

import "fmt"

func maxVal(arr []int) float64 {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return float64(max)
}

func Radixsort(arr []int) []int {
	max := maxVal(arr)
	exp := 1
	for int(max/exp) > 0 {
		arr = countSort(arr, exp)
		exp *= 10
	}
	return arr
}

func countSort(arr []int, exp int) []int {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		index := arr[i] / exp
		count[index%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		index := arr[i] / exp
		output[count[index%10]-1] = arr[i]
		count[index%10]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}

	return arr
}

func main() {
	array := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println(Radixsort(array))
}
