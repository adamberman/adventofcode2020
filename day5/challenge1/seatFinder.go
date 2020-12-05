package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func rowToBinary(row string) string {
	var bRow [7]string
	for i, val := range row {
		if val == 'F' {
			bRow[i] = "0"
		} else {
			bRow[i] = "1"
		}
	}
	return strings.Join(bRow[:], "")
}

func columnToBinary(column string) string {
	var bColumn [3]string
	for i, val := range column {
		if val == 'L' {
			bColumn[i] = "0"
		} else {
			bColumn[i] = "1"
		}
	}
	return strings.Join(bColumn[:], "")
}

func findSeatID(seat string) int {
	row := seat[:7]
	column := seat[7:]
	rowAsBinary := rowToBinary(row)
	columnAsBinary := columnToBinary(column)
	rowAsDec, err := strconv.ParseInt(rowAsBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	columnAsDec, err := strconv.ParseInt(columnAsBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(rowAsDec*8 + columnAsDec)
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
	var maxID int
	for scanner.Scan() {
		id := findSeatID(scanner.Text())
		if id > maxID {
			maxID = id
		}
	}
	fmt.Println(maxID)
}
