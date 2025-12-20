package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("**** Welcome to URL Shortner ****")
	fmt.Println("Would you like to proceed?")
	fmt.Println("1:Yes")
	fmt.Println("2:No")
	userInput := bufio.NewReader(os.Stdin)
	enteredValue, err := userInput.ReadString('\n')
	userSelection_one := strings.TrimSpace(enteredValue)
	fmt.Println("Heyy", enteredValue)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	//Create a map that will store the values of short -> long
	maps := make(map[string]string)
	switch userSelection_one {
	case "1":
		fmt.Println("Please enter the URL that you would like to be shortened:  ")
		userURL := bufio.NewReader(os.Stdin)
		urlValue, err := userURL.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading URL:", err)
			return
		}
		// Generate a short code of random string for the string.
		// Example: Generate a random 6-character string
		charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		length := 6
		replacementShortURL := generateRandomString(charset, length)

		fmt.Println("Here is Shortened URL code (it's saved in memory):", replacementShortURL)
		maps[replacementShortURL] = urlValue
	case "2":
		fmt.Println("See You Next Time!")
		return
	}
}

// When function is not dependent on any hidden state we use regular function(no receiver)
func generateRandomString(charset string, length int) string {
	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			panic(err)
		}
		sb.WriteByte(charset[int(b[0])%len(charset)])
	}
	return sb.String()
}
