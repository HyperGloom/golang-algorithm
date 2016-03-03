/*
 * https://en.wikipedia.org/wiki/Karatsuba_algorithm
 */
package main
import (
	"fmt"
	"math"
)

func getDecimalDigits(num int64) uint  {
	var result uint

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

func getHighAndLowDigits(num int64, digits uint) (int64, int64) {
	divisor := int64(math.Pow(10, float64(digits)))

	if num >= divisor {
		return num / divisor, num % divisor
	} else {
		return 0, num
	}
}

func karatsuba(x int64, y int64) int64 {
	var max_digits uint
	positive := true

	if x == 0 || y == 0 {
		return 0
	}

	if (x > 0 && y < 0) || (x < 0 && y > 0) {
		positive = false
	}

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	if x < 10 || y < 10 {
		return  x * y
	}

	x_digits := getDecimalDigits(x)
	y_digits := getDecimalDigits(y)

	if x_digits >= y_digits {
		max_digits = x_digits / 2;
	} else {
		max_digits = y_digits / 2;
	}

	x_high, x_low := getHighAndLowDigits(x, max_digits)
	y_high, y_low := getHighAndLowDigits(y, max_digits)

	z0 := karatsuba(x_low, y_low)
	z1 := karatsuba((x_low + x_high),(y_low + y_high))
	z2 := karatsuba(x_high, y_high)

	if positive {
		return (z2 * int64(math.Pow(10, float64(2 * max_digits)))) + (z1 - z2 - z0) * int64(math.Pow(10, float64(max_digits))) + z0
	} else {
		return -((z2 * int64(math.Pow(10, float64(2 * max_digits)))) + (z1 - z2 - z0) * int64(math.Pow(10, float64(max_digits))) + z0)
	}

}

func main() {
	fmt.Println(karatsuba(1234, 5678))
	fmt.Println(karatsuba(1234, -5678))
}
