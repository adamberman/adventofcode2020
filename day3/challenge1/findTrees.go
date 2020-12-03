package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	var m []string
	for scanner.Scan() {
		m = append(m, scanner.Text())
	}
	var x int
	var trees int
	for y := 0; y < len(m); y++ {
		if m[y][x] == '#' {
			trees++
		}
		x = (x + 3) % len(m[y])
	}
	fmt.Println(trees)
}
