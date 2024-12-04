package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	x int
	y int
}

type Pattern struct {
	directions []Direction
	symbols    []string
}

func day4() {

	var grid [140][140]string

	if data, err := os.ReadFile("./input4.txt"); err == nil {
		lines := strings.Split(string(data), "\n")
		for row, line := range lines {
			if line != "" {
				for col, char := range line {
					grid[row][col] = string(char)
				}
			}
		}
	}

	directions := []Direction{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	targetWord := "XMAS"

	xpatterns := []Pattern{
		{[]Direction{{0, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}, []string{"A", "M", "M", "S", "S"}},
		{[]Direction{{0, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}, []string{"A", "M", "S", "S", "M"}},
		{[]Direction{{0, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}, []string{"A", "S", "M", "M", "S"}},
		{[]Direction{{0, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}, []string{"A", "S", "S", "M", "M"}},
	}

	getGrid := func(row, col int) string {
		if row < 0 || col < 0 || row >= 140 || col >= 140 {
			return "-"
		}
		return grid[row][col]
	}

	total1 := 0
	total2 := 0

	for row := range grid {
		for col := range grid[row] {
			for _, dir := range directions {
				word := ""
				for step := range len(targetWord) {
					word += getGrid(row+dir.x*step, col+dir.y*step)
				}
				if word == targetWord {
					total1++
				}
			}
			for _, pattern := range xpatterns {
				valid := true
				for i, direction := range pattern.directions {
					if getGrid(row+direction.x, col+direction.y) != pattern.symbols[i] {
						valid = false
					}
				}
				if valid {
					total2++
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)

}
