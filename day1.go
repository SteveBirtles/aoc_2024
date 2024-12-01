package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func day1() {

	if data, err := os.ReadFile("./input1.txt"); err == nil {

		var list1 []int
		var list2 []int

		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if left, right, valid := strings.Cut(line, " "); valid {
				list1 = append(list1, ToInt(left))
				list2 = append(list2, ToInt(right))
			}
		}

		slices.Sort(list1)
		slices.Sort(list2)

		total1 := 0
		for i := range list1 {
			total1 += AbsInt(list1[i] - list2[i])
		}

		fmt.Printf("Part 1: %d\n", total1)

		total2 := 0
		for i := range list1 {
			count := 0
			for j := range list2 {
				if list1[i] == list2[j] {
					count++
				}
			}
			total2 += list1[i] * count
		}

		fmt.Printf("Part 2: %d\n", total2)

	}

}
