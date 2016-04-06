package main

import "fmt"

func fibonacci(n int) uint {
	m := make([]uint, n+1)

	m[0] = 0
	m[1] = 1
	if n == 0 || n == 1 {
		return m[n]
	}

	for i := 2; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n]
}

func main() {
	fmt.Println(fibonacci(100))
}
