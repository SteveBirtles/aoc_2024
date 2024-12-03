package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func ToInt(s string) int {
	s = strings.TrimSpace(s)
	if value, err := strconv.Atoi(s); err == nil {
		return value
	} else {
		fmt.Printf("Unable to convert '%s' to int! Returning 0...\n", s)
		return 0
	}
}
