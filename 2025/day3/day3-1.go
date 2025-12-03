package main

import (
	"bufio"
	"fmt"
	"os"
)

func processBank(bank []int) int {

	joltage := 0

	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			tmp := bank[i]*10 + bank[j]
			if tmp > joltage {
				joltage = tmp
			}
		}
	}

	fmt.Printf("Max Bank joltage = %d\n", joltage)

	return joltage
}

func main() {

	filename := os.Args[1]
	totalJoltage := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	fmt.Println("Starting")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var bank []int
		for _, char := range line {
			bank = append(bank, int(char-'0'))
		}
		fmt.Println(bank)
		joltage := processBank(bank)
		totalJoltage += joltage
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Printf("totalJoltage: %d\n", totalJoltage)
}
