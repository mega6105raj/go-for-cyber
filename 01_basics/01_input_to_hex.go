package main

import (
	"fmt"	//Used for printing output and reading input
	"encoding/hex"	//This is for hexadecimal encoding
)

func main() {
	var input string	//Declare a variable to hold user input
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	byteData := []byte(input)	//Convert string to byte slice
	hexOutput := hex.EncodeToString(byteData)	//Encode byte slice into hex format

	fmt.Println("Hex representation:", hexOutput)	//Print the hex result
}
