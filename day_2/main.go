package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var allGames []string

func parseGameId(line string, games *map[string]string) {
	splitString := strings.Split(line, ":")
	gameSection := splitString[0]

	gameSectionSplit := strings.Split(gameSection, " ")[1]
	(*games)[gameSectionSplit] = splitString[1]
	allGames = append(allGames, gameSectionSplit)
}

func removeDuplicates(slice []string) []string {
	// Create a map to store unique elements
	seen := make(map[string]bool)
	result := []string{}

	// Loop through the slice, adding elements to the map if they haven't been seen before
	for _, val := range slice {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("There was an error opening input.txt")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// red, green, blue := 12, 13, 14
	games := &map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		parseGameId(line, games)
		// allGames = append(allGames, (*games)[0])
	}

	fmt.Println(allGames)

	lowestPerKey := map[string]int{}

	// Format: 1: "1 red, 4 blue; 4 red, 5 green"
	for key, value := range *games {
		splitSectionsPerGame := strings.Split(value, ";")

		lowestGreen, lowestRed, lowestBlue := 0, 0, 0

		for _, section := range splitSectionsPerGame {
			rgbValues := strings.Split(section, ", ")

			fmt.Println(key)

			for _, rgbVal := range rgbValues {
				rgbVal = strings.TrimSpace(rgbVal)
				fmt.Println(rgbVal)

				num := strings.Split(rgbVal, " ")[0]
				color := strings.Split(rgbVal, " ")[1]
				numOfCube, err := strconv.Atoi(string(num))

				if err != nil {
					log.Fatal(err)
				}

				switch true {
				case color[0] == 'r':
					fmt.Println("Red - ", numOfCube)
					if numOfCube > lowestRed {
						lowestRed = numOfCube
					}
				case color[0] == 'g':
					fmt.Println("Green", numOfCube)
					if numOfCube > lowestGreen {
						lowestGreen = numOfCube
					}
				case color[0] == 'b':
					fmt.Println("Blue", numOfCube)
					if numOfCube > lowestBlue {
						lowestBlue = numOfCube
					}
				}
			}
			fmt.Println("====")
		}

		fmt.Println(lowestRed, " - ", lowestGreen, " - ", lowestBlue)
		lowestPerKey[key] = lowestRed * lowestGreen * lowestBlue
		// fmt.Println(key, " ", value, " -- ", splitSectionsPerGame)
	}

	fmt.Println(lowestPerKey)

	totes := 0
	for _, val := range lowestPerKey {
		totes += val
	}

	fmt.Println("TOTES", totes)
	// fmt.Println("FINAL ", total)
}
