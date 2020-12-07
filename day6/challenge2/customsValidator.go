package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processGroup(group []string) int {
	fmt.Println(group)
	elemsCount := make(map[string]int)
	for _, row := range group {
		for i := 0; i < len(row); i++ {
			elemsCount[string(row[i])]++
		}
	}
	groupLength := len(group)
	var count int
	for k := range elemsCount {
		if elemsCount[k] == groupLength {
			count++
		}
	}
	return count
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
	group := []string{}
	var count int
	for scanner.Scan() {
		row := scanner.Text()
		fmt.Println(row)
		if len(row) == 0 {
			count += processGroup(group)
			group = nil
		} else {
			group = append(group, row)
		}
	}
	count += processGroup(group)
	fmt.Println(count)
}
