package main

import (
	"fmt"
	"os"
	"strings"
)

const MAP_SIZE = 130

func day6() {

	var worldMap [MAP_SIZE][MAP_SIZE]int
	x := -1
	y := -1

	if data, err := os.ReadFile("./input6.txt"); err == nil {
		lines := strings.Split(string(data), "\n")
		for row, line := range lines {
			if line != "" {
				for col, char := range line {
					if char == '#' {
						worldMap[col][row] = -1
					} else if char == '^' {
						x = col
						y = row
					}
				}
			}
		}
	}

	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{-1, 0, 1, 0}
	dir := 0

	worldMap[x][y] = 1
	for x+dx[dir] >= 0 && y+dy[dir] >= 0 && x+dx[dir] < MAP_SIZE && y+dy[dir] < MAP_SIZE {
		if worldMap[x+dx[dir]][y+dy[dir]] == -1 {
			dir = (dir + 1) % 4
		} else {
			x += dx[dir]
			y += dy[dir]
			worldMap[x][y] = 1
		}
	}

	total1 := 0

	for row := range MAP_SIZE {
		for col := range MAP_SIZE {
			if worldMap[col][row] == 1 {
				total1++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", total1)

}
