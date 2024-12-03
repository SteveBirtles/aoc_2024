package main

import (
	"fmt"
	"os"
	"strings"
)

type TokenType int

const (
	Start TokenType = iota
	Comma
	End
	Int
	Unknown
	Do
	Dont
)

type Token struct {
	raw       string
	tokenType TokenType
}

var TokenStrings = []Token{
	{"mul(", Start},
	{",", Comma},
	{")", End},
	{"do()", Do},
	{"don't()", Dont},
}

func day3() {

	if data, err := os.ReadFile("./input3.txt"); err == nil {

		lines := append(strings.Split(string(data), "\n"), " ")
		tokens := []Token{}
		candidateToken := ""

		for _, line := range lines {

			for _, char := range line {

				nextChar := string(char)

				if IsInt(candidateToken) && !IsInt(candidateToken+nextChar) {
					tokens = append(tokens, Token{candidateToken, Int})
					candidateToken = ""
				}

				for _, token := range TokenStrings {
					if len(candidateToken) >= len(token.raw) && candidateToken[len(candidateToken)-len(token.raw):] == token.raw {
						if len(candidateToken) > len(token.raw) {
							tokens = append(tokens, Token{candidateToken[:len(candidateToken)-len(token.raw)], Unknown})
						}
						tokens = append(tokens, Token{token.raw, token.tokenType})
						candidateToken = ""
						break
					}
				}

				candidateToken += nextChar

			}

		}

		total1 := 0
		total2 := 0
		part2active := true
		t := 0
		for t < len(tokens)-4 {
			if tokens[t].tokenType == Do {
				part2active = true
			} else if tokens[t].tokenType == Dont {
				part2active = false
			} else if tokens[t].tokenType == Start &&
				tokens[t+1].tokenType == Int &&
				tokens[t+2].tokenType == Comma &&
				tokens[t+3].tokenType == Int &&
				tokens[t+4].tokenType == End {
				total1 += ToInt(tokens[t+1].raw) * ToInt(tokens[t+3].raw)
				if part2active {
					total2 += ToInt(tokens[t+1].raw) * ToInt(tokens[t+3].raw)
				}
				t += 5
				continue
			}
			t++
		}

		fmt.Printf("Part 1: %d\n", total1)
		fmt.Printf("Part 2: %d\n", total2)

	}

}
