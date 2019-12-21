package main

import (
	"fmt"
	"strconv"
	"strings"
)

type doubleGroupMatcher func(map[rune]int) bool

func main() {
	input := "165432-707912"
	fmt.Printf("Valid password of range %s with at least one double: %d\n", input, CountValidPasswords(input, withAtLeastOneDouble))
	fmt.Printf("Valid password of range %s with exactly one double: %d\n", input, CountValidPasswords(input, withExactlyOneDouble))
}

// CountValidPasswords find the secret given the range
func CountValidPasswords(guess string, checkDouble doubleGroupMatcher) int {
	parts := strings.Split(guess, "-")
	begin, end := toInt(parts[0]), toInt(parts[1])

	validPasswords := 0

	for password := begin; password <= end; password++ {
		if isValidPassword(strconv.FormatInt(int64(password), 10), checkDouble) {
			validPasswords++
		}
	}
	return validPasswords
}

func isValidPassword(password string, checkDouble doubleGroupMatcher) bool {
	if isDecreasing(password) || !checkDouble(getDoubledDigits(password)) {
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

func getDoubledDigits(password string) map[rune]int {
	doubled := make(map[rune]int, 6)
	for _, l := range password {
		doubled[l]++
	}
	return doubled
}

func withAtLeastOneDouble(doubled map[rune]int) bool {
	for l := range doubled {
		if doubled[l] > 1 {
			return true
		}
	}
	return false
}

func withExactlyOneDouble(doubled map[rune]int) bool {
	for l := range doubled {
		if doubled[l] == 2 {
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
