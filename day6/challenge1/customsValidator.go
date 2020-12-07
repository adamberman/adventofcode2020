package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processGroup(group []string) int {
	fmt.Println(group)
	var count int
	uniqueElems := make(map[string]bool)
	for _, row := range group {
		for i := 0; i < len(row); i++ {
			if !uniqueElems[string(row[i])] {
				uniqueElems[string(row[i])] = true
				count++
			}
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
