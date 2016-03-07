/*
 * https://en.wikipedia.org/wiki/Merge_sort
 */
package main

import "fmt"

func mergeSort(slice []int) []int {

	var i int

	if len(slice) <= 1 {
		return slice
	}

	ret := make([]int, len(slice))

	l := mergeSort(slice[0:(len(slice) / 2)])
	r := mergeSort(slice[(len(slice) / 2):])

	for i = 0; i < len(slice); i++ {

		if len(l) <= 0 || len(r) <= 0 {
			break
		}
		if l[0] <= r[0] {
			ret[i] = l[0]
			l = l[1:]
		} else {
			ret[i] = r[0]
			r = r[1:]
		}
	}

	if len(l) > 0 {
		copy(ret[i:], l)
	}

	if len(r) > 0 {
		copy(ret[i:], r)
	}

	return ret
}

func main() {
	array := []int{10, 9, 7, 4, 34, 76, 5, 9, 3, 28}
	//array := []int{10}
	array = mergeSort(array)
	fmt.Println(array)
}
