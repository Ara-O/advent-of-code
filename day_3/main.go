package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseFileInto2dArray(mappedSchemes [][]string) {
	for i := range mappedSchemes {
		mappedSchemes[i] = make([]string, 155)
	}

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		for i, letter := range line {
			mappedSchemes[row][i] = string(letter)
		}

		row++
	}
}

type Key struct {
	number         string
	hasAdjacentKey bool
	direction      string
}

func main() {
	schemaKeys := "$%&/*+@-#="
	mappedSchemes := make([][]string, 140)
	var keys []Key

	parseFileInto2dArray(mappedSchemes)

	for i, row := range mappedSchemes {
		currentNumber, direction := "", ""
		hasAdjacentKey := false
		for j, _ := range row {
			letter := mappedSchemes[i][j]
			if num, err := strconv.Atoi(letter); err == nil {
				currentNumber += strconv.Itoa(num)

				// up and down
				if i < 139 {
					if strings.ContainsAny(schemaKeys, mappedSchemes[i+1][j]) {
						hasAdjacentKey = true
						direction = "vertical"
					}
				}

				if i > 0 {
					if strings.ContainsAny(schemaKeys, mappedSchemes[i-1][j]) {
						hasAdjacentKey = true
						direction = "vertical"
					}
				}

				// right and left
				if j > 0 {
					if strings.ContainsAny(schemaKeys, mappedSchemes[i][j-1]) {
						hasAdjacentKey = true
						direction = "horizontal"
					}
				}

				if j < 139 {
					if strings.ContainsAny(schemaKeys, mappedSchemes[i][j+1]) {
						hasAdjacentKey = true
						direction = "horizontal"
					}
				}

				//diagonals
				if i > 0 && j > 0 {
					if strings.ContainsAny(schemaKeys, mappedSchemes[i-1][j-1]) {
						hasAdjacentKey = true
						direction = "top left diagonal"
					}
				}

				if i > 0 && j < 139 {
					if strings.ContainsAny(schemaKeys, mappedSchemes[i-1][j+1]) {
						hasAdjacentKey = true
						direction = "top right diagonal"
					}
				}

				if i < 139 && j < 139 {
					if strings.ContainsAny(schemaKeys, mappedSchemes[i+1][j+1]) {
						hasAdjacentKey = true
						direction = "bottom right diagonal"
					}
				}

				if i < 139 && j > 0 {
					if strings.ContainsAny(schemaKeys, mappedSchemes[i+1][j-1]) {
						hasAdjacentKey = true
						direction = "bottom left diagonal"
					}
				}

			} else {
				if len(currentNumber) > 0 {
					fmt.Println(currentNumber)
					keys = append(keys, Key{
						number:         currentNumber,
						hasAdjacentKey: hasAdjacentKey,
						direction:      direction,
					})
					currentNumber = ""
					hasAdjacentKey = false
					direction = ""
				}
			}
		}
	}

	// fmt.Printf("%+v", keys)

	sum := 0
	for _, i := range keys {
		fmt.Println(i)
		if i.hasAdjacentKey == true {
			if num, err := strconv.Atoi(i.number); err == nil {
				sum += num
			}
		}
	}

	fmt.Println("TOTAL SUM: ", sum)
}
