package main

import (
	"fmt"
	"os"
	"strconv"
)

type steps struct {
	count int
	s     [][]int
}

func calcSteps(num int) *steps {
	m := make(map[int]*steps)

	m[0] = &steps{}
	m[0].count = 0
	m[0].s = make([][]int, m[0].count)

	m[1] = &steps{}
	m[1].count = 1
	m[1].s = make([][]int, m[1].count)
	m[1].s[0] = make([]int, 0)
	m[1].s[0] = append(m[1].s[0], 1)

	m[2] = &steps{}
	m[2].count = 2
	m[2].s = make([][]int, m[2].count)
	m[2].s[0] = make([]int, 0)
	m[2].s[0] = append(m[2].s[0], 1)
	m[2].s[0] = append(m[2].s[0], 1)
	m[2].s[1] = make([]int, 0)
	m[2].s[1] = append(m[2].s[1], 2)

	m[3] = &steps{}
	m[3].count = 4
	m[3].s = make([][]int, m[3].count)
	m[3].s[0] = make([]int, 0)
	m[3].s[0] = append(m[3].s[0], 1)
	m[3].s[0] = append(m[3].s[0], 1)
	m[3].s[0] = append(m[3].s[0], 1)
	m[3].s[1] = make([]int, 0)
	m[3].s[1] = append(m[3].s[1], 1)
	m[3].s[1] = append(m[3].s[1], 2)
	m[3].s[2] = make([]int, 0)
	m[3].s[2] = append(m[3].s[2], 2)
	m[3].s[2] = append(m[3].s[2], 1)
	m[3].s[3] = make([]int, 0)
	m[3].s[3] = append(m[3].s[3], 3)

	for i := 4; i <= num; i++ {
		fmt.Println(i)
		//var j int
		m[i] = &steps{}
		m[i].count = m[i-1].count + m[i-2].count + m[i-3].count
		/* m[i].s = make([][]int, m[i].count)

		for k := 0; k < m[i - 1].count; k++ {
			m[i].s[j] = make([]int, 0)
			m[i].s[j] = append(m[i].s[j], 1)
			m[i].s[j] = append(m[i].s[j], m[i - 1].s[k]...)
			j++
		}

		for k := 0; k < m[i - 2].count; k++ {
			m[i].s[j] = make([]int, 0)
			m[i].s[j] = append(m[i].s[j], 2)
			m[i].s[j] = append(m[i].s[j], m[i - 2].s[k]...)
			j++
		}

		for k := 0; k < m[i - 3].count; k++ {
			m[i].s[j] = make([]int, 0)
			m[i].s[j] = append(m[i].s[j], 3)
			m[i].s[j] = append(m[i].s[j], m[i - 3].s[k]...)
			j++
		}
		delete(m, i - 3) */
	}

	return m[num]
}
func main() {
	num, _ := strconv.Atoi(os.Args[1])
	result := calcSteps(num)
	fmt.Println(result.count)
	/* for i := 0; i < result.count; i++ {
		fmt.Println(result.s[i])
	} */
}
