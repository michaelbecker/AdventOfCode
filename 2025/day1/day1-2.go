package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func rotLeft(dial, distance int) (int, int) {
	zeros := 0

	for distance >= 100 {
		zeros++
		fmt.Println("Zero")
		distance -= 100
	}

	originalDial := dial
	dial = dial - distance
	if dial == 0 {
		zeros++
		fmt.Println("Zero")
	} else if dial < 0 {
		dial += 100
		if originalDial != 0 {
			zeros++
		}
		fmt.Println("Zero")
	}

	return zeros, dial
}

func rotRight(dial, distance int) (int, int) {
	zeros := 0

	for distance >= 100 {
		zeros++
		fmt.Println("Zero")
		distance -= 100
	}

	dial = dial + distance
	if dial == 0 {
		zeros++
		fmt.Println("Zero")
	} else if dial > 99 {
		dial -= 100
		zeros++
		fmt.Println("Zero")
	}

	return zeros, dial
}

func main() {

	var dial int = 50
	filename := os.Args[1]
	countZeros := 0

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
		dir := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Printf("Executing: %c - %d - Dial: %d\n", dir, distance, dial)

		switch dir {

		case 'L':
			var zeros int
			zeros, dial = rotLeft(dial, distance)
			countZeros += zeros
		case 'R':
			var zeros int
			zeros, dial = rotRight(dial, distance)
			countZeros += zeros

		default:
			fmt.Printf("Unknown direction: %c", dir)
			return
		}

		fmt.Printf("Dial: %d\n", dial)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Printf("Password: %d\n", countZeros)
}
