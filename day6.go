package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

const MAP_SIZE = 130

func day6() {

	var worldMap [MAP_SIZE][MAP_SIZE]int
	startX := -1
	startY := -1

	if data, err := os.ReadFile("./input6.txt"); err == nil {
		lines := strings.Split(string(data), "\n")
		for row, line := range lines {
			if line != "" {
				for col, char := range line {
					if char == '#' {
						worldMap[col][row] = -1
					} else if char == '^' {
						startX = col
						startY = row
					}
				}
			}
		}
	}

	x := startX
	y := startY
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{-1, 0, 1, 0}
	dir := 0

	var turningPoints []Coordinate
	var possiblities []Coordinate

	for x+dx[dir] >= 0 && y+dy[dir] >= 0 && x+dx[dir] < MAP_SIZE && y+dy[dir] < MAP_SIZE {
		if worldMap[x+dx[dir]][y+dy[dir]] == -1 {
			turningPoints = append(turningPoints, Coordinate{x, y})
			dir = (dir + 1) % 4
		} else {
			x += dx[dir]
			y += dy[dir]
			if worldMap[x][y] == 0 {
				worldMap[x][y] = 10 + dir
			}
			if len(turningPoints) >= 3 {
				corners := []Coordinate{
					turningPoints[len(turningPoints)-3],
					turningPoints[len(turningPoints)-2],
					turningPoints[len(turningPoints)-1],
					{x, y}}
				topLeft := Coordinate{MAP_SIZE, MAP_SIZE}
				bottomRight := Coordinate{-1, -1}
				for _, corner := range corners {
					topLeft.x = min(topLeft.x, corner.x)
					topLeft.y = min(topLeft.y, corner.y)
					bottomRight.x = max(bottomRight.x, corner.x)
					bottomRight.y = max(bottomRight.y, corner.y)
				}
				isRect := true
				for _, corner := range corners {
					if corner.x != topLeft.x && corner.x != bottomRight.x ||
						corner.y != topLeft.y && corner.y != bottomRight.y {
						isRect = false
						break
					}
				}
				if isRect {
					rectClear := true
					for i := topLeft.x; i <= bottomRight.x; i++ {
						if worldMap[i][topLeft.y] < 0 || worldMap[i][bottomRight.y] < 0 {
							rectClear = false
						}
					}
					for j := topLeft.y; j <= bottomRight.y; j++ {
						if worldMap[topLeft.x][j] < 0 || worldMap[bottomRight.x][j] < 0 {
							rectClear = false
						}
					}
					if rectClear {
						possiblities = append(possiblities, Coordinate{x + dx[dir], y + dy[dir]})
						worldMap[turningPoints[len(turningPoints)-3].x][turningPoints[len(turningPoints)-3].y] = 2
						worldMap[turningPoints[len(turningPoints)-2].x][turningPoints[len(turningPoints)-2].y] = 2
						worldMap[turningPoints[len(turningPoints)-1].x][turningPoints[len(turningPoints)-1].y] = 2
						worldMap[x][y] = 2
						worldMap[x+dx[dir]][y+dy[dir]] = 3
					} else {
						worldMap[x+dx[dir]][y+dy[dir]] = 4
					}
				}
			}
		}

	}

	total1 := 0

	for row := range MAP_SIZE {
		for col := range MAP_SIZE {
			if worldMap[col][row] > 0 {
				total1++
			}
			if worldMap[col][row] == 10 {
				fmt.Print("^")
			} else if worldMap[col][row] == 11 {
				fmt.Print(">")
			} else if worldMap[col][row] == 12 {
				fmt.Print("v")
			} else if worldMap[col][row] == 13 {
				fmt.Print("<")
			} else if worldMap[col][row] == 2 {
				fmt.Print("╬")
			} else if worldMap[col][row] == 3 {
				fmt.Print("█")
			} else if worldMap[col][row] == 4 {
				fmt.Print("X")
			} else if worldMap[col][row] == -1 {
				fmt.Print("▒")
			} else if worldMap[col][row] == 0 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", len(possiblities)) //total2)

}
