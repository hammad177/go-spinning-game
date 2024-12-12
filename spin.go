package main

import (
	"fmt"
	"math/rand"
)

func GenerateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}

	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func GetSpin(reel []string, rows int, cols int) [][]string {
	result := make([][]string, rows)

	for col := 0; col < cols; col++ {
		selectedIndex := map[int]bool{}
		for row := 0; row < rows; row++ {
			for true {
				radomIndex := getRandomNumber(0, len(reel)-1)
				_, exists := selectedIndex[radomIndex]
				if !exists {
					result[row] = append(result[row], reel[radomIndex])
					selectedIndex[radomIndex] = true
					break
				}
			}
		}
	}

	return result
}

func PrintSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf("%s ", symbol)
			if j < len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}
