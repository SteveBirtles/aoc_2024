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

func day8() {

	antennas := map[string][]Coordinate{}

	if data, err := os.ReadFile("./input8.txt"); err == nil {
		lines := strings.Split(string(data), "\n")
		for y, line := range lines {
			if line != "" {
				for x, char := range line {
					if char != '.' {
						label := string(char)
						if group, found := antennas[label]; !found {
							antennas[label] = []Coordinate{{x, y}}
						} else {
							antennas[label] = append(group, Coordinate{x, y})
						}
					}
				}
			}
		}
	}

	var isAntinode1 [50][50]bool
	var isAntinode2 [50][50]bool
	total1 := 0
	total2 := 0

	for _, group := range antennas {
		for i, first := range group {
			for j, second := range group {
				if i == j {
					continue
				}
				dx := second.x - first.x
				dy := second.y - first.y

				position := Coordinate{second.x, second.y}

				for count := 0; !(position.x < 0 || position.y < 0 || position.x >= 50 || position.y >= 50); count++ {

					if count == 1 {
						if !isAntinode1[position.x][position.y] {
							total1++
						}
						isAntinode1[position.x][position.y] = true
					}

					if !isAntinode2[position.x][position.y] {
						total2++
					}
					isAntinode2[position.x][position.y] = true

					position.x += dx
					position.y += dy
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)

}
