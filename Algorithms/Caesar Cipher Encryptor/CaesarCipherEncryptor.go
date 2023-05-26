package main

import "fmt"

func main() {
	var (
		input string
		num   int
	)

	fmt.Print("Enter a string: ")
	fmt.Scan(&input)
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	result := caesarCipher(input, num)
	fmt.Println("\nResult: ", result)
}
