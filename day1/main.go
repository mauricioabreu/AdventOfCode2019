package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	masses := strings.Split(string(fileContent), ("\n"))
	requiredFuel := 0
	requiredFuelWithFueld := 0
	for _, mass := range masses {
		i, err := strconv.Atoi(mass)
		if err != nil {
			continue
		}
		requiredFuel = requiredFuel + getFuel(i)
		requiredFuelWithFueld += getFuelWithFuel(i)
	}
	fmt.Printf("Fuel required to load all modules: %d\n", requiredFuel)
	fmt.Printf("Fuel required to load all modules and all the fuel: %d\n", requiredFuelWithFueld)
}

func getFuel(mass int) int {
	return mass/3 - 2
}

func getFuelWithFuel(mass int) int {
	requiredFuel := getFuel(mass)
	for fuel := getFuel(requiredFuel); fuel > 0; fuel = getFuel(fuel) {
		requiredFuel += fuel
	}
	return requiredFuel
}
