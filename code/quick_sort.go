package main

import "fmt"

func quickSort(array []int) []int {

	if len(array) == 1 {
		return array
	}

	left := 0
	right := len(array) - 1
	pivot := array[0]

	for right > left {
		for ; right > left; right-- {
			if pivot > array[right] {
				array[left] = array[right]
				break
			}
		}
		for ; right > left; left++ {
			if pivot < array[left] {
				array[right] = array[left]
				break
			}
		}
	}

	array[right] = pivot

	if right > 1 {
		quickSort(array[0:right])
	}

	if right < len(array)-1 {
		quickSort(array[(right + 1):])
	}

	return array
}

func main() {
	array := []int{6, 4, 7, 3, 2, 1, 8, 5}
	fmt.Println(quickSort(array))
	array = []int{6, 4, 7, 9, 3, 2, 1, 8, 5}
	fmt.Println(quickSort(array))
	array = []int{6, 4, 7, 9, 3, 2, 1, 8, 5, 3}
	fmt.Println(quickSort(array))
	array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(quickSort(array))
	array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(quickSort(array))
}
