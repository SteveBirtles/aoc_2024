package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type OrderRule struct {
	before int
	after  int
}

var orderRules []OrderRule

// Found a better way whilst doing part 2...
/* func updateOK(update []int) bool {
	for i := 1; i < len(update); i++ {
		for _, orderRule := range orderRules {
			if update[i] == orderRule.before {
				for j := 0; j < i; j++ {
					if update[j] == orderRule.after {
						return false
					}
				}
			}
		}
	}
	return true
} */

func orderSortFunc(a, b int) int {
	for _, orderRule := range orderRules {
		if orderRule.after == a && orderRule.before == b {
			return 1
		} else if orderRule.after == b && orderRule.before == a {
			return -1
		}
	}
	return 0
}

func day5() {

	var updates [][]int

	if data, err := os.ReadFile("./input5.txt"); err == nil {
		lines := strings.Split(string(data), "\n")
		section := 1
		for _, line := range lines {
			if line == "" {
				section = 2
			} else if section == 1 {
				before, after, _ := strings.Cut(line, "|")
				orderRules = append(orderRules, OrderRule{ToInt(before), ToInt(after)})
			} else if section == 2 {
				update := []int{}
				for _, item := range strings.Split(line, ",") {
					update = append(update, ToInt(item))
				}
				updates = append(updates, update)
			}
		}
	}

	total1 := 0
	total2 := 0

	for _, update := range updates {
		//if updateOK(update) {   // Found a better way whilst doing part 2
		if slices.IsSortedFunc(update, orderSortFunc) {
			total1 += update[len(update)/2]
		} else {
			slices.SortFunc(update, orderSortFunc)
			total2 += update[len(update)/2]
		}
	}

	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)

}
