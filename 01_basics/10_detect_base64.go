/* Documentation Section:
	This program detects whether a given file contains valid Base64-encoded content.
	- It reads the file
	- Checks character pattern (A-Z, a-z, 0-9, +, /, =)
	- Attempts to decode using base64.StdEncoding
	Id decoding succeeds without error, it's likely a valid Base64-encoded file */

package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var filePath string
	fmt.Print("Enter path to file: ")
	fmt.Scanln(&filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text()
	}

	//Remove whitespace/newlines
	cleaned := regexp.MustCompile(`\s+`).ReplaceAllString(content, "")

	//Check if it matches base64 character set
	base64Pattern := regexp.MustCompile(`^[A-Za-z0-9+/=]+$`)
	if !base64Pattern.MatchString(cleaned) {
		fmt.Println("File is NOT Base64 encoded (invalid characters)")
		return
	}

	//Try decoding
	_, err = base64.StdEncoding.DecodeString(cleaned)
	if err != nil {
		fmt.Println("File is NOT valid Base64 (decode error):", err)
	} else {
		fmt.Println("File is likely Base64 encoded.")
	}
}
