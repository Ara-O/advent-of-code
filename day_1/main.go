package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	mappedDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	if err != nil {
		log.Fatal("There was an error opening input.txt")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalSum := 0

	for scanner.Scan() {
		// digitsFound := []int{}
		line := scanner.Text()

		firstDigit, lastDigit := "", ""
		newLine := ""

		// First digit
		for _, l := range line {
			newLine += string(l)

			for key := range mappedDigits {
				if strings.Contains(newLine, key) {
					// fmt.Println("Found ", newLine, " ", key)
					firstDigit = key
					break
				}
			}

			if len(firstDigit) > 0 {
				break
			}
		}

		//Last digit
		newLine = ""

		for i := len(line) - 1; i >= 0; i-- {
			letter := string(line[i])

			newLine = letter + newLine
			for key := range mappedDigits {
				if strings.Contains(newLine, key) {
					lastDigit = key
					break
				}
			}

			if len(lastDigit) > 0 {
				break
			}
		}

		combined := mappedDigits[firstDigit] + mappedDigits[lastDigit]

		if num, err := strconv.Atoi(combined); err == nil {
			totalSum += num
		}
	}

	print("TOTAL: ", totalSum)
}
