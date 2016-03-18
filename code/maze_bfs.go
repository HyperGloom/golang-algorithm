package main

import (
	"fmt"
)

type coordinate struct {
	x, y int
}

type node struct {
	c    coordinate
	step int
}

func coordinateIsValid(maze [][]int, c coordinate) bool {
	if c.x >= 0 && c.x <= len(maze)-1 && c.y >= 0 && c.y <= len(maze[c.x])-1 && maze[c.x][c.y] == 0 {
		return true
	} else {
		return false
	}
}

func bfs(maze [][]int, start coordinate, end coordinate) (int, bool) {

	if !coordinateIsValid(maze, start) || !coordinateIsValid(maze, end) {
		return 0, false
	}

	var head *node
	var queue []*node
	hasVisited := make(map[coordinate]bool)

	queue = append(queue, &node{start, 1})
	hasVisited[start] = true

	for len(queue) > 0 {

		head = queue[0]
		queue = queue[1:]

		if head.c.x == end.x && head.c.y == end.y {
			return head.step, true
		}

		if coordinateIsValid(maze, coordinate{head.c.x, head.c.y - 1}) && !hasVisited[coordinate{head.c.x, head.c.y - 1}] {
			hasVisited[coordinate{head.c.x, head.c.y - 1}] = true
			queue = append(queue, &node{coordinate{head.c.x, head.c.y - 1}, head.step + 1})
		}

		if coordinateIsValid(maze, coordinate{head.c.x, head.c.y + 1}) && !hasVisited[coordinate{head.c.x, head.c.y + 1}] {
			hasVisited[coordinate{head.c.x, head.c.y + 1}] = true
			queue = append(queue, &node{coordinate{head.c.x, head.c.y + 1}, head.step + 1})
		}

		if coordinateIsValid(maze, coordinate{head.c.x - 1, head.c.y}) && !hasVisited[coordinate{head.c.x - 1, head.c.y}] {
			hasVisited[coordinate{head.c.x - 1, head.c.y}] = true
			queue = append(queue, &node{coordinate{head.c.x - 1, head.c.y}, head.step + 1})
		}

		if coordinateIsValid(maze, coordinate{head.c.x + 1, head.c.y}) && !hasVisited[coordinate{head.c.x + 1, head.c.y}] {
			hasVisited[coordinate{head.c.x + 1, head.c.y}] = true
			queue = append(queue, &node{coordinate{head.c.x + 1, head.c.y}, head.step + 1})
		}

	}
	return head.step, false
}

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0, 1},
		{0, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{1, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 1}}
	step, err := bfs(maze, coordinate{0, 0}, coordinate{2, 3})
	fmt.Printf("%v:%v\n", err, step)
}
