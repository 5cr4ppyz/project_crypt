package main // Make this an executable package

import (
	"fmt" // Format for output and input
)

var price float64 = 95000.50
var quantity float64 = 0.05

func main() {
	symbol := "BTC-USDT"
	
	computer_msg := "Computer: Yo scrappyz, whats up?"
	scrappyz_msg := fmt.Sprintf("Scrappyz: %s is trading at %v, I'm about to buy %v worth", symbol, price, quantity) 

	fmt.Println(computer_msg)
	fmt.Println(scrappyz_msg)

	fmt.Println("\n--- Initializing ---")
	//var total_cost float64 = price * quantity
	//fmt.Println(total_cost); // Couldn't assign 0.05 as interger

	// Bru you cannot assign leverage as float the complier denies
	//var leverage = float64(5) // Turn the value type to Float
	
	//var total_pos = leverage * price
	//fmt.Println(total_pos)

	// Compiler Yell
 	//var compiler_yell = ".main.go:24:18: invalid operation: leverage * price (mismatched types int and float64)";
	
	const BYBIT_API_URL = "https://api.bybit.com"
	const max_entries = 3

	fmt.Printf("Fetching data from %s, with a %d max attempt", BYBIT_API_URL, max_entries)






}