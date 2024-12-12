package main

import (
	"fmt"
)

func checkWin(spin [][]string, multiplier map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		is_win := true

		for i, symbol := range row[1:] {
			if symbol != row[i] {
				lines = append(lines, 0)
				is_win = false
				break
			}
		}

		if is_win {
			lines = append(lines, multiplier[row[0]])
		}
	}

	return lines
}

func main() {
	symbols := map[string]uint{
		"A": 5,
		"B": 8,
		"C": 10,
		"D": 13,
	}
	multiplier := map[string]uint{
		"A": 15,
		"B": 10,
		"C": 8,
		"D": 5,
	}

	symbolArr := GenerateSymbolArray(symbols)

	balance := uint(200)
	GetName()

	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		spin := GetSpin(symbolArr, 3, 3)
		fmt.Println("Spinning the reels...")
		PrintSpin(spin)

		lines := checkWin(spin, multiplier)

		for i, line := range lines {
			if line > 0 {
				fmt.Printf("Congratulations: You won $%d (%dx) on line %d\n", line*bet, line, i+1)
				balance += line * bet
			}
		}
	}

	fmt.Printf("Your balance is now $%d\n", balance)
}
