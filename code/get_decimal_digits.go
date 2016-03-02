/*
 * Get the number of digits of an integer.
 */
package main
import (
	"fmt"
)

func getDecimalDigits(num int64) int  {
	var result int

	if num == 0 {
		return 1
	}

	if num < 0 {
		num = -num
	}
	for num > 0 {
		result++
		num = num / 10
	}

	return result
}

func main() {
	fmt.Println(getDecimalDigits(0))
	fmt.Println(getDecimalDigits(9))
	fmt.Println(getDecimalDigits(-9))
	fmt.Println(getDecimalDigits(20))
	fmt.Println(getDecimalDigits(-20))
	fmt.Println(getDecimalDigits(101))
	fmt.Println(getDecimalDigits(-101))
}
