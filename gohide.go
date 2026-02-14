package main

import (
	"fmt"
)

// Clear Screen
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Check if this is the first time the user is running the program
func isFirstTime() bool {

}

// Display main menu
func displayMainMenu() {
	clearScreen()

	fmt.Println("\n\t+-----------+")
	fmt.Println("\t|  Go Hide  |")
	fmt.Println("\t+-----------+")
	fmt.Println("\n >>  Zero-Trust Password Manager <<")

	fmt.Println("\n [1] Login")
	fmt.Println(" [2] Register")
	fmt.Println(" [3] Exit")

	var choice int
	fmt.Print("\nEnter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Login selected")
	case 2:
		fmt.Println("Register selected")
	case 3:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice, please try again.")
	}

}

// Function to write data into files

func main() {

	// var fname, lname string

	// fmt.Print("Enter your first name: ")
	// fmt.Scan(&fname)
	// fmt.Print("Enter your last name: ")
	// fmt.Scan(&lname)

	// fmt.Println("Hello", fname, lname)

	// Display Main Menu
	displayMainMenu()

}
