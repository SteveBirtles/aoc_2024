package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var stones []int

func day11() {

	if data, err := os.ReadFile("./input11.txt"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, rawValue := range strings.Split(lines[0], " ") {
			stones = append(stones, ToInt(rawValue))
		}
	}

	fmt.Println(stones)

	for range 5 {

		for i := 0; i < len(stones); {

			stoneString := fmt.Sprintf("%d", stones[i])
			n := len(stoneString)

			if stones[i] == 0 {
				stones[i] = 1
				i++
			} else if n%2 == 0 {
				newStones := []int{ToInt(stoneString[:n/2]), ToInt(stoneString[n/2:])}
				stones = slices.Delete(stones, i, i+1)
				stones = slices.Insert(stones, i, newStones...)
				i += 2
			} else {
				stones[i] *= 2048
				i++
			}

		}

		fmt.Println(stones)
		fmt.Println(len(stones))

	}

}
