package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "165432-707912"
	fmt.Printf("Valid password of range %s: %d\n", input, CountValidPasswords(input))
}

// CountValidPasswords find the secret given the range
func CountValidPasswords(guess string) int {
	parts := strings.Split(guess, "-")
	begin, end := toInt(parts[0]), toInt(parts[1])

	validPasswords := 0

	for password := begin; password <= end; password++ {
		if isValidPassword(strconv.FormatInt(int64(password), 10)) {
			validPasswords++
		}
	}
	return validPasswords
}

func isValidPassword(password string) bool {
	if isDecreasing(password) || !hasDouble(password) {
		return false
	}
	return true
}

func isDecreasing(password string) bool {
	if password[0] > password[1] ||
		password[1] > password[2] ||
		password[2] > password[3] ||
		password[3] > password[4] ||
		password[4] > password[5] {
		return true
	}
	return false
}

func hasDouble(password string) bool {
	doubled := make(map[rune]int, 6)
	for _, l := range password {
		doubled[l]++
		if doubled[l] > 1 {
			return true
		}
	}
	return false
}

func toInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return number
}
