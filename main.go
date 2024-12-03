package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 0 {
		fmt.Println("Please specify a day to run as the first argument.")
		return
	}

	day := ToInt(os.Args[1])

	switch day {
	case 1:
		day1()
	case 2:
		day2()
	case 3:
		day3()
	default:
		fmt.Printf("Day %d not implemented.\n", day)
	}

}
