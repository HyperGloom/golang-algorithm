package main

import (
	"fmt"
)

const maxInt = int(^uint(0) >> 1)

type vertex byte

type edge struct {
	v   vertex
	len int
}

type path struct {
	dis  int
	v    []vertex
	find bool
}

func dijkstra(m map[vertex][]edge, ver []vertex, start, end vertex) *path {
	p := make(map[vertex]*path)

	for _, v := range ver {
		p[v] = &path{dis: maxInt}
	}

	p[start].dis = 0
	p[start].find = true
	lastMinVertex := start
	for {
		if lastMinVertex == end {
			return p[lastMinVertex]
		}

		for _, e := range m[lastMinVertex] {
			if p[lastMinVertex].dis+e.len < p[e.v].dis {
				p[e.v].dis = p[lastMinVertex].dis + e.len
				p[e.v].v = append(p[lastMinVertex].v, e.v)
				fmt.Println(e.v, p[e.v].v)
			}
		}

		minDist := maxInt
		minVertex := lastMinVertex
		for k, v := range p {
			if !v.find && v.dis < minDist {
				minVertex = k
				minDist = v.dis
			}
		}
		if minVertex == lastMinVertex {
			return nil
		} else {
			p[minVertex].find = true
			lastMinVertex = minVertex
		}
	}
}

func main() {
	m := make(map[vertex][]edge)

	m['a'] = append(m['a'], edge{'b', 7})
	m['a'] = append(m['a'], edge{'c', 9})
	m['a'] = append(m['a'], edge{'f', 14})

	m['b'] = append(m['b'], edge{'c', 10})
	m['b'] = append(m['b'], edge{'d', 15})

	m['c'] = append(m['c'], edge{'d', 11})
	m['c'] = append(m['c'], edge{'f', 2})

	m['d'] = append(m['d'], edge{'e', 6})

	m['e'] = append(m['e'], edge{'f', 9})

	/* for i, v := range m {
		fmt.Println(i, v)
	} */

	fmt.Println(dijkstra(m, []vertex{'a', 'b', 'c', 'd', 'e', 'f'}, 'a', 'e'))
}
