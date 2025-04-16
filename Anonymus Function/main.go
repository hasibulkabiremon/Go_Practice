package main

import "fmt"

func printReports(intro, body, outro string) {
	// ?
	printCostReport(func(s string) int{
		return 2*len(s)
	},intro)

	printCostReport(func(s string) int{
		return 3*len(s)
	},body)

	printCostReport(func(s string) int{
		return 4*len(s)
	},outro)
}

// don't touch below this line

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
