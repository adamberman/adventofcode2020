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
	positions := strings.Split(split[0], "-")
	pos1, err := strconv.Atoi(positions[0])
	if err != nil {
		log.Fatal(err)
	}
	pos2, err := strconv.Atoi(positions[1])
	if err != nil {
		log.Fatal(err)
	}
	character := split[1][0]
	password := split[2]
	if (password[pos1-1] == character && password[pos2-1] != character) || (password[pos1-1] != character && password[pos2-1] == character) {
		return 1
	}

	return 0
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
