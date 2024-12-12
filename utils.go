package main

import "fmt"

func GetName() string {
	var name string

	fmt.Println("Welcome to Tim's Casino Game: ")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, to Tim's Casino Game!, Let's Play \n", name)

	return name
}

func GetBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		_, err := fmt.Scan(&bet)
		if err != nil || bet < 0 {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if bet > balance {
			fmt.Println("You cannot bet more than your balance!")
		} else {
			break
		}
	}

	return bet
}
