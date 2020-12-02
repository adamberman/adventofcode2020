package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func validatePasswordLine(line string) int {
	split := strings.Split(line, " ")
	counts := strings.Split(split[0], "-")
	min, err := strconv.Atoi(counts[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(counts[1])
	if err != nil {
		log.Fatal(err)
	}
	character := split[1][0]
	password := split[2]
	var charCount int
	for i := range password {
		if password[i] == character {
			charCount++
		}
	}

	if charCount > max || charCount < min {
		return 0
	} else {
		return 1
	}
}

func main() {
	var filePath string
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	} else {
		log.Fatal("Please provide a file path as an argument")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() {
		count += validatePasswordLine(scanner.Text())
	}

	fmt.Println(count)
}
