package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
			dial = dial - distance
			if dial < 0 {
				for {
					dial += 100
					if dial >= 0 {
						break
					}
				}
			}
		case 'R':
			dial = dial + distance
			if dial > 99 {
				for {
					dial -= 100
					if dial <= 99 {
						break
					}
				}
			}
		default:
			fmt.Printf("Unknown direction: %c", dir)
			return
		}

		if dial == 0 {
			countZeros++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Printf("Password: %d\n", countZeros)
}
