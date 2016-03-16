package main

import (
	"fmt"
)

func dfs(array []int, book map[int]bool, result []int, step int) {
	if step == len(array) {
		fmt.Printf("% v\n", result)
		return
	}

	for i := 0; i < len(array); i++ {
		if !book[array[i]] {
			book[array[i]] = true
			result[step] = array[i]
			dfs(array, book, result, step+1)
			book[array[i]] = false
		}
	}
}

func permutation(array []int) {
	book := make(map[int]bool)
	result := make([]int, len(array))

	dfs(array, book, result, 0)
}

func main() {
	array := []int{1, 2, 3, 4}
	permutation(array)
}
