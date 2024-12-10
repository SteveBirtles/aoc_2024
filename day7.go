package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day7() {

	total1 := 0
	total2 := 0

	if data, err := os.ReadFile("./input7.txt"); err == nil {

		lines := strings.Split(string(data), "\n")

		for _, line := range lines {

			if line == "" {
				break
			}
			rawResult, rawParts, _ := strings.Cut(line, ":")
			result := ToInt(rawResult)

			_ = result

			parts := []int{}
			for _, part := range strings.Split(rawParts[1:], " ") {
				parts = append(parts, ToInt(part))
			}

		loop1:
			for i := range int(math.Pow(2, float64(len(parts)-1))) {

				//output := fmt.Sprintf("%d = %d", result, parts[0])

				pattern := "000000000000" + strconv.FormatInt(int64(i), 2)
				pattern = pattern[len(pattern)-(len(parts)-1):]
				subTotal := parts[0]
				for i, bit := range pattern {
					if bit == '0' {
						subTotal += parts[i+1]
						//output += fmt.Sprintf(" + %d", parts[i+1])
					} else {
						subTotal *= parts[i+1]
						//output += fmt.Sprintf(" * %d", parts[i+1])
					}
				}
				if subTotal == result {
					total1 += result
					break loop1
				}
			}

		loop2:
			for i := range int(math.Pow(3, float64(len(parts)-1))) {

				//output := fmt.Sprintf("%d = %d", result, parts[0])

				pattern := "000000000000" + strconv.FormatInt(int64(i), 3)

				pattern = pattern[len(pattern)-(len(parts)-1):]
				subTotal := parts[0]
				for i, bit := range pattern {
					if bit == '0' {
						subTotal += parts[i+1]
						//output += fmt.Sprintf(" + %d", parts[i+1])
					} else if bit == '1' {
						subTotal *= parts[i+1]
						//output += fmt.Sprintf(" * %d", parts[i+1])
					} else {
						concat := fmt.Sprintf("%d%d", subTotal, parts[i+1])
						subTotal = ToInt(concat)
						//output += fmt.Sprintf(" || %d", parts[i+1])
					}
				}
				if subTotal == result {
					total2 += result
					break loop2
				}
			}

		}

	}

	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)

}
