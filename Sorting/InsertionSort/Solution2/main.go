package main

import "fmt"

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 && array[j] < array[j-1] {
			array[j], array[j-1] = array[j-1], array[j]
			j--
		}
	}
	return array
}

func main() {
	fmt.Println(InsertionSort([]int{8, 5, 2, 9, 5, 6, 3}))
}
