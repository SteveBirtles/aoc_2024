package main

import (
	"fmt"
	"os"
	"strings"
)

const GRID_SIZE = 57

var grid [GRID_SIZE][GRID_SIZE]int

func traverse1(x, y, searchValue int) int {
	if x < 0 || y < 0 || x >= GRID_SIZE || y >= GRID_SIZE {
		return 0
	} else if grid[x][y] == searchValue {
		if searchValue == 9 {
			grid[x][y] = -1
			return 1
		} else {
			return traverse1(x+1, y, searchValue+1) +
				traverse1(x, y+1, searchValue+1) +
				traverse1(x-1, y, searchValue+1) +
				traverse1(x, y-1, searchValue+1)
		}
	} else {
		return 0
	}
}

func traverse2(x, y, searchValue int) int {
	if x < 0 || y < 0 || x >= GRID_SIZE || y >= GRID_SIZE {
		return 0
	} else if grid[x][y] == searchValue {
		if searchValue == 9 {
			return 1
		} else {
			return traverse2(x+1, y, searchValue+1) +
				traverse2(x, y+1, searchValue+1) +
				traverse2(x-1, y, searchValue+1) +
				traverse2(x, y-1, searchValue+1)
		}
	} else {
		return 0
	}
}

func day10() {

	if data, err := os.ReadFile("./input10.txt"); err == nil {
		lines := strings.Split(string(data), "\n")
		for row, line := range lines {
			if line != "" {
				for col, char := range line {
					grid[row][col] = ToInt(string(char))
				}
			}
		}
	}

	total1 := 0
	total2 := 0
	for x := range GRID_SIZE {
		for y := range GRID_SIZE {
			if grid[x][y] == 0 {
				value1 := traverse1(x, y, 0)
				for i := range GRID_SIZE {
					for j := range GRID_SIZE {
						if grid[i][j] == -1 {
							grid[i][j] = 9
						}
					}
				}
				//fmt.Printf("(%d, %d) --> %d\n", x, y, value)
				total1 += value1

				value2 := traverse2(x, y, 0)
				fmt.Printf("(%d, %d) --> %d\n", x, y, value2)
				total2 += value2

			}
		}
	}
	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)

}
