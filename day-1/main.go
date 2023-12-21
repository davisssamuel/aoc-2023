package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var wordToNum = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func checkPair(pair [2]string, str string) (newPair [2]string) {
	newPair = pair
	if pair[0] != "" {
		newPair[1] = str
		return
	} else {
		newPair[0], newPair[1] = str, str
		return
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		var pair [2]string

		for i := 0; i < len(line); i++ {

			// check characters
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				pair = checkPair(pair, string(line[i]))
			}

			// check substrings
			for j := i + 1; j <= len(line); j++ {
				n := wordToNum[string(line[i:j])]
				if n != "" {
					pair = checkPair(pair, n)
					break
				}
			}
		}
		num, _ := strconv.Atoi(pair[0] + pair[1])
		println(num)
		sum = sum + num
	}
	println("sum:", sum)
}
