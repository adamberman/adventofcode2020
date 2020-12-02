package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	var arr []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, val)
	}
	for i := range arr {
		for j := range arr {
			if i == j {
				continue
			}
			if arr[i]+arr[j] == 2020 {
				fmt.Println(arr[i] * arr[j])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
