package main

import (
	"fmt"
	"os"
)

func main() {

	day := 0
	if len(os.Args) >= 2 {
		day = ToInt(os.Args[1])
	}

	switch day {
	case 0:
		fmt.Println("Day 1...")
		day1()
		fmt.Println("Day 2...")
		day2()
		fmt.Println("Day 3...")
		day3()
		fmt.Println("Day 4...")
		day4()
		fmt.Println("Day 5...")
		day5()
		fmt.Println("Day 6...")
		day6()
		fmt.Println("Day 7...")
		day7()
		fmt.Println("Day 8...")
		day8()
		fmt.Println("Day 9...")
		day9()
		fmt.Println("Day 10...")
		day10()
		fmt.Println("Day 11...")
		day11()
	case 1:
		day1()
	case 2:
		day2()
	case 3:
		day3()
	case 4:
		day4()
	case 5:
		day5()
	case 6:
		day6()
	case 7:
		day7()
	case 8:
		day8()
	case 9:
		day9()
	case 10:
		day10()
	case 11:
		day11()
	default:
		fmt.Printf("Day %d not implemented.\n", day)
	}

}
