package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	idSum := 0
	powSum := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		possibleGame := true

		red := 0
		green := 0
		blue := 0

		for i := 2; i < len(line); i++ {
			s := string(line[i])

			// remove last char (punctuation) on all words except the last
			if i < (len(line) - 1) {
				s = s[:(len(s) - 1)]
			}

			switch s {
			case "red":
				n, _ := strconv.Atoi(string(line[i-1]))
				if n > 12 {
					possibleGame = false
				}
				if n > red {
					red = n
				}
			case "green":
				n, _ := strconv.Atoi(string(line[i-1]))
				if n > 13 {
					possibleGame = false
				}
				if n > green {
					green = n
				}
			case "blue":
				n, _ := strconv.Atoi(string(line[i-1]))
				if n > 14 {
					possibleGame = false
				}
				if n > blue {
					blue = n
				}
			}
		}

		// check if game was possible
		if possibleGame {
			s := string(line[1])
			n, _ := strconv.Atoi(s[:(len(s) - 1)])
			idSum = idSum + n
		}
		powSum = powSum + (red * green * blue)
	}
	println("ID sum:", idSum)
	println("Power sum:", powSum)
}
