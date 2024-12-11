package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func day9() {

	hardDrive0 := []int{}

	if data, err := os.ReadFile("./input9.txt"); err == nil {

		flipFlop := true
		fileId := 0

		lines := strings.Split(string(data), "\n")
		for _, rawDigit := range lines[0] {
			if flipFlop {
				fileSize := ToInt(string(rawDigit))
				for range fileSize {
					hardDrive0 = append(hardDrive0, fileId)
				}
				fileId++
			} else {
				freeSpaceSize := ToInt(string(rawDigit))
				for range freeSpaceSize {
					hardDrive0 = append(hardDrive0, -1)
				}
			}
			flipFlop = !flipFlop
		}

		updateFreeSpacePointer := func(hardDrive []int, start int) int {
			for i := start + 1; i < len(hardDrive); i++ {
				if hardDrive[i] == -1 {
					return i
				}
			}
			panic("No free space!")
		}

		contiguousFreeSpace := func(hardDrive []int, start int) int {
			for i := start; i < len(hardDrive); i++ {
				if hardDrive[i] != -1 {
					return i - start
				}
			}
			return len(hardDrive) - 1 - start
		}

		establishFileSize := func(hardDrive []int, start int) int {
			id := hardDrive[start]
			for i := start + 1; i < len(hardDrive); i++ {
				if hardDrive[i] != id {
					return i - start
				}
			}
			return len(hardDrive) - 1 - start
		}

		freeSpacePointer := -1
		hardDrive1 := slices.Clone(hardDrive0)

	firstLoop:
		for i := len(hardDrive1) - 1; i >= 0; i-- {
			if hardDrive1[i] == -1 {
				continue
			}
			freeSpacePointer = updateFreeSpacePointer(hardDrive1, freeSpacePointer)
			if freeSpacePointer > i {
				break firstLoop
			}
			hardDrive1[freeSpacePointer] = hardDrive1[i]
			hardDrive1[i] = -1
		}

		total1 := 0

		for i, value := range hardDrive1 {
			if value == -1 {
				continue
			}
			total1 += i * value
		}

		fmt.Printf("Part 1: %d\n", total1)

		freeSpacePointer = -1
		hardDrive2 := slices.Clone(hardDrive0)

		currentFileId := -1

	secondLoop:
		for i := len(hardDrive2) - 1; i >= 0; i-- {
			if hardDrive2[i] == -1 || hardDrive2[i] == currentFileId {
				continue
			} else if currentFileId != -1 {
				currentFileSize := establishFileSize(hardDrive2, currentFileId)

				for {
					freeSpacePointer = updateFreeSpacePointer(hardDrive2, freeSpacePointer)
					if freeSpacePointer > i {
						break secondLoop
					}
					if contiguousFreeSpace(hardDrive2, freeSpacePointer) >= currentFileSize {
						break
					}
				}

				for j := 0; j < currentFileSize; j++ {
					hardDrive2[freeSpacePointer+j] = hardDrive2[i+j]
					hardDrive2[i+j] = -1
				}
			}

		}

		total2 := 0

		for i, value := range hardDrive2 {
			if value == -1 {
				continue
			}
			total2 += i * value
		}

		fmt.Printf("Part 2: %d\n", total2)

	}
}
