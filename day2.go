package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func checkValues(values []string) bool {

	increasing := ToInt(values[1])-ToInt(values[0]) > 0

	for i := range values {
		if i == 0 {
			continue
		}
		diff := ToInt(values[i]) - ToInt(values[i-1])
		if diff == 0 || diff > 3 || diff < -3 ||
			(diff < 0 && diff >= -3 && increasing) ||
			(diff > 0 && diff <= 3 && !increasing) {
			return false
		}
	}

	return true
}

func day2() {

	if data, err := os.ReadFile("./input2.txt"); err == nil {

		lines := strings.Split(string(data), "\n")
		count1 := 0
		count2 := 0

		for _, line := range lines {
			values := strings.Split(line, " ")
			if checkValues(values) {
				count1++
				count2++
			} else {
				for j := range values {
					if checkValues(slices.Delete(slices.Clone(values), j, j+1)) {
						count2++
						break
					}
				}
			}
		}

		fmt.Printf("Part 1: %d\n", count1)
		fmt.Printf("Part 2: %d\n", count2)

	}

}
