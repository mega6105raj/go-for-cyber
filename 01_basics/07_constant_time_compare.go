/* Documentation Section :
	When comparing sensitive values (liek passwords, HMACs, or API keys), regular comparison (==) may leak sensitive information
Timing attacks exploit this to gradually guess values based on how long comparisons take.
Go provides subtle.ConstantTimeCompare() to mitigae this risk */

package main

import (
	"crypto/subtle"
	"fmt"
)

func main() {
	var a, b string
	fmt.Print("Enter first string: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second string : ")
	fmt.Scanln(&b)

	//Convert to byte slices
	aBytes := []byte(a)
	bBytes := []byte(b)

	//Compare lengths first
	if len(aBytes) != len(bBytes) {
		fmt.Println("Strings do not match (different lengths)")
		return
	}

	//Constant-time comparison
	if subtle.ConstantTimeCompare(aBytes, bBytes) == 1 {
		fmt.Println("Strings match (constant-time check)")
	} else {
		fmt.Println("Strings do NOT match (constant-tiem check)")
	}
}
