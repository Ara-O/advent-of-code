package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("There was an error opening input.txt")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalSum := 0

	for scanner.Scan() {
		digitsFound := []int{}
		line := scanner.Text()

		for _, l := range line {
			letter := string(l)

			//If an error wasn't thrown when converting to a number
			if num, err := strconv.Atoi(letter); err == nil {
				digitsFound = append(digitsFound, num)
			}
		}
		if len(digitsFound) == 1 {
			digitsFound = append(digitsFound, digitsFound[0])
		}

		sum := strconv.Itoa(digitsFound[0]) + strconv.Itoa(digitsFound[len(digitsFound)-1])

		if num, err := strconv.Atoi(sum); err == nil {
			totalSum += num
		}
	}

	print("TOTAL: ", totalSum)
}
