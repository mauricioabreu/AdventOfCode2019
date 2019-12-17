package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

// IntCode process a program of numbers and return
// the first position of the generated sequence
func IntCode(program []int) int {
	pos := 0

	for {
		opCode := program[pos]

		inputA := program[pos+1]
		inputB := program[pos+2]
		outputPos := program[pos+3]

		switch opCode {
		case 1:
			program[outputPos] = program[inputA] + program[inputB]
		case 2:
			program[outputPos] = program[inputA] * program[inputB]
		case 99:
			return program[0]
		default:
			return 0
		}
		pos += 4
	}
}

func restore(sequence []int, noun int, verb int) []int {
	restoredProgram := make([]int, len(sequence))
	copy(restoredProgram, sequence)
	restoredProgram[1], restoredProgram[2] = noun, verb
	return restoredProgram
}

func main() {
	sequence := readSequence("input.txt")
	program := restore(sequence, 12, 2)
	fmt.Printf("Intcode part 1: %d\n", IntCode(program))

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program = restore(sequence, noun, verb)
			result := IntCode(program)
			if result == 19690720 {
				fmt.Printf("IntCode part 2: %d\n", 100*noun+verb)
				os.Exit(0)
			}
		}
	}
}
