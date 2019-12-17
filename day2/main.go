package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readSequence(file string) []int {
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sequence := []int{}

	numbers := strings.Split(string(fileContent), (","))
	for _, number := range numbers {
		sequence = append(sequence, toInt(number))
	}
	return sequence
}

func toInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return number
}

// IntCode process a sequence of numbers and return
// the first position of the generated sequence
func IntCode(sequence []int) int {
	restoredProgram := restore(sequence)

	pos := 0

	for {
		opCode := restoredProgram[pos]

		inputA := restoredProgram[pos+1]
		inputB := restoredProgram[pos+2]
		outputPos := restoredProgram[pos+3]

		switch opCode {
		case 1:
			restoredProgram[outputPos] = restoredProgram[inputA] + restoredProgram[inputB]
		case 2:
			restoredProgram[outputPos] = restoredProgram[inputA] * restoredProgram[inputB]
		case 99:
			return restoredProgram[0]
		default:
			return 0
		}
		pos += 4
	}
}

func restore(sequence []int) []int {
	restoredProgram := make([]int, len(sequence))
	copy(restoredProgram, sequence)
	restoredProgram[1], restoredProgram[2] = 12, 2
	return restoredProgram
}

func main() {
	sequence := readSequence("input.txt")
	fmt.Printf("Intcode is: %d\n", IntCode(sequence))
}
