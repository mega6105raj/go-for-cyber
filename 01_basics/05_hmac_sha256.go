/* Documentation Section:
	This program computes the HMAC (Hash-based Message Authentication Code) of an input message using a secret key
and the SHA-256 hash function
HMACs are essential for:
	- Verifying integrity + authenticit of messages
	- Used in JWTs, TLS, API key signing
	- Avoiding tampering (unlike plain hashes)
 Unlike SHA-256 alone, HMAC uses a secret key, making it resistan to length extension and collision attacks
*/

package main

import (
	"crypto/hmac" 	//For HMAC computaion
	"crypto/sha256" 	//For SHA-256 hash function
	"encoding/hex" //To display HMAC in hex
	"fmt"
)

func main() {
	var message string
	var key string

	fmt.Print("Enter message : ")
	fmt.Scanln(&message)

	fmt.Print("Enter key : ")
	fmt.Scanln(&key)

	h := hmac.New(sha256.New, []byte(key))	//Create new HMAC using SHA-256 and key
	h.Write([]byte(message))	//Input the message to hash

	hmacValue := h.Sum(nil)	//Ge the final HMAC as byte slice
	hmacHex := hex.EncodeToString(hmacValue)

	fmt.Println("HMAC-SHA256 (hex):", hmacHex)
}
