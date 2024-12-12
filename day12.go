package main

import (
	"fmt"
	"os"
	"strings"
)

const GARDEN_SIZE = 140

type GardenPlot struct {
	plantType int
	groupId   int
}

var garden [GARDEN_SIZE][GARDEN_SIZE]GardenPlot

func day12() {

	if data, err := os.ReadFile("./input12.txt"); err == nil {
		id := 0
		lines := strings.Split(string(data), "\n")
		for row, line := range lines {
			if line != "" {
				for col, char := range line {
					garden[row][col] = GardenPlot{int(char), id}
					id++
				}
			}
		}
	}

	unifyGroups := func(source, target int) {
		for i := 0; i < GARDEN_SIZE; i++ {
			for j := 0; j < GARDEN_SIZE; j++ {
				if garden[i][j].groupId == source {
					garden[i][j].groupId = target
				}
			}
		}
	}

	for x := 0; x < GARDEN_SIZE; x++ {
		for y := 0; y < GARDEN_SIZE; y++ {
			if x < GARDEN_SIZE-1 && garden[x+1][y].plantType == garden[x][y].plantType {
				unifyGroups(garden[x+1][y].groupId, garden[x][y].groupId)
			}
			if y < GARDEN_SIZE-1 && garden[x][y+1].plantType == garden[x][y].plantType {
				unifyGroups(garden[x][y+1].groupId, garden[x][y].groupId)
			}
		}
	}

	groupSize := make(map[int]int)
	groupPerimeter := make(map[int]int)

	for x := 0; x < GARDEN_SIZE; x++ {
		for y := 0; y < GARDEN_SIZE; y++ {

			//fmt.Printf("%d\t", garden[x][y].groupId)

			id := garden[x][y].groupId
			groupSize[id] = groupSize[id] + 1

			pSoFar := groupPerimeter[id]
			if x == 0 || garden[x-1][y].groupId != id {
				pSoFar++
			}
			if y == 0 || garden[x][y-1].groupId != id {
				pSoFar++
			}
			if x == GARDEN_SIZE-1 || garden[x+1][y].groupId != id {
				pSoFar++
			}
			if y == GARDEN_SIZE-1 || garden[x][y+1].groupId != id {
				pSoFar++
			}
			groupPerimeter[id] = pSoFar

		}
		//fmt.Println()
	}

	total1 := 0
	for k, v := range groupSize {

		p := groupPerimeter[k]

		//fmt.Printf("%d: %d x %d = %d\n", k, v, p, v*p)

		total1 += v * p
	}

	fmt.Printf("Part 1: %d\n", total1)
}
