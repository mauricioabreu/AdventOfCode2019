package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	masses := strings.Split(string(fileContent), ("\n"))
	sum := 0
	for _, mass := range masses {
		i, err := strconv.Atoi(mass)
		if err != nil {
			continue
		}
		sum = sum + getFuel(i)
	}
	fmt.Printf("Sum required to load all modules: %d\n", sum)
}

func getFuel(mass int) int {
	return int(math.Floor((float64(mass) / 3) - 2))
}
