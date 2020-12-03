package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findTrees(slopeX int, slopeY int, m []string) int {
	var x int
	var trees int

	for y := 0; y < len(m); y += slopeY {
		if m[y][x] == '#' {
			trees++
		}
		x = (x + slopeX) % len(m[y])
	}
	return trees
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
	var m []string
	for scanner.Scan() {
		m = append(m, scanner.Text())
	}
	fmt.Println(findTrees(1, 1, m) * findTrees(3, 1, m) * findTrees(5, 1, m) * findTrees(7, 1, m) * findTrees(1, 2, m))
}
