package main

import (
	"fmt"
	"os"
	"strings"
)

var dx = [4]int{0, 1, 0, -1}
var dy = [4]int{-1, 0, 1, 0}

const MAP_SIZE = 130

type WorldCell struct {
	obstruction bool
	visited     bool
	directions  [4]bool
}

var worldMap [MAP_SIZE][MAP_SIZE]WorldCell

func day6() {

	startX := -1
	startY := -1
	dir := 0

	if data, err := os.ReadFile("./input6.txt"); err == nil {
		lines := strings.Split(string(data), "\n")
		for row, line := range lines {
			if line != "" {
				for col, char := range line {
					if char == '#' {
						worldMap[col][row].obstruction = true
					} else if char == '^' {
						startX = col
						startY = row
					}
				}
			}
		}
	}

	total1 := 0
	total2 := 0
	x := startX
	y := startY

	for x+dx[dir] >= 0 && y+dy[dir] >= 0 && x+dx[dir] < MAP_SIZE && y+dy[dir] < MAP_SIZE {

		if worldMap[x+dx[dir]][y+dy[dir]].obstruction {
			dir = (dir + 1) % 4
		}

		if !worldMap[x+dx[dir]][y+dy[dir]].visited && !(x+dx[dir] == startX && y+dy[dir] == startY) {

			worldMap2 := worldMap
			worldMap2[x+dx[dir]][y+dy[dir]].obstruction = true

			x2 := x
			y2 := y
			dir2 := dir

			for x2+dx[dir2] >= 0 && y2+dy[dir2] >= 0 && x2+dx[dir2] < MAP_SIZE && y2+dy[dir2] < MAP_SIZE {
				if worldMap2[x2+dx[dir2]][y2+dy[dir2]].obstruction {
					dir2 = (dir2 + 1) % 4
				} else {
					x2 += dx[dir2]
					y2 += dy[dir2]
					if worldMap2[x2][y2].visited && worldMap2[x2][y2].directions[dir2] {
						total2++
						break
					}
					worldMap2[x2][y2].visited = true
					worldMap2[x2][y2].directions[dir2] = true
				}
			}
		}

		x += dx[dir]
		y += dy[dir]
		if !worldMap[x][y].visited {
			total1++
		}
		worldMap[x][y].visited = true
		worldMap[x][y].directions[dir] = true
	}

	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)

}
