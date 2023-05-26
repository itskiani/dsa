package main

import "fmt"

func caesarCipher(str string, num int) string {
	shift, offset := int32(num%26), int32(26)
	letters := []int32(str)

	for i, char := range letters {
		if char >= 'a' && char+shift <= 'z' {
			char += shift
		} else {
			char += shift - offset
		}

		letters[i] = char
	}

	return string(letters)
}

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
