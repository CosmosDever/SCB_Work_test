package main

import (
	"fmt"
	"unicode"
)

func minimumNumber(n int32, password string) int32 {
	// Function to check if the password contains at least one digit
	containsDigit := func(password string) bool {
		for _, char := range password {
			if unicode.IsDigit(char) {
				return true
			}
		}
		return false
	}

	// Function to check if the password contains at least one lowercase letter
	containsLowercase := func(password string) bool {
		for _, char := range password {
			if unicode.IsLower(char) {
				return true
			}
		}
		return false
	}

	// Function to check if the password contains at least one uppercase letter
	containsUppercase := func(password string) bool {
		for _, char := range password {
			if unicode.IsUpper(char) {
				return true
			}
		}
		return false
	}

	// Function to check if the password contains at least one special character
	containsSpecialCharacter := func(password string) bool {
		for _, char := range password {
			if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
				return true
			}
		}
		return false
	}

	// Return the number of characters to make the password strong
	// Initialize the count of characters to add to 0
	var count int32 = 0

	// Check if the password contains at least one digit
	if !containsDigit(password) {
		count++
	}

	// Check if the password contains at least one lowercase letter
	if !containsLowercase(password) {
		count++
	}

	// Check if the password contains at least one uppercase letter
	if !containsUppercase(password) {
		count++
	}

	// Check if the password contains at least one special character
	if !containsSpecialCharacter(password) {
		count++
	}

	// Check if the password is less than 6 characters
	if int(n)+int(count) < 6 {
		count += int32(6 - (int(n) + int(count)))
	}

	return count
}

func main() {
	// Test cases
	// Test case 1
	n1 := int32(3)
	password1 := "Ab1"
	// Expected output: 3
	fmt.Println(minimumNumber(n1, password1))

	// Test case 2
	n2 := int32(11)
	password2 := "#HackerRank"

	// Expected output: 1
	fmt.Println(minimumNumber(n2, password2))
}
