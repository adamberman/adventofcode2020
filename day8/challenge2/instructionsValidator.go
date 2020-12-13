package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInstruction(instruction string) (string, int) {
	instArr := strings.Split(instruction, " ")
	op := instArr[0]
	count, err := strconv.Atoi(instArr[1][1:])
	if err != nil {
		log.Fatal(err)
	}
	if string(instArr[1][0]) == "-" {
		return op, (0 - count)
	}

	return op, count
}

func findInfiniteLoop(instructions []string) (bool, int) {
	executed := make(map[int]bool)
	var accum int
	i := 0
	for {
		if i == len(instructions) {
			return true, accum
		}
		_, ok := executed[i]
		if ok {
			return false, accum
		}
		executed[i] = true
		operation, value := parseInstruction(instructions[i])
		if operation == "acc" {
			accum += value
			i++
		} else if operation == "nop" {
			i++
		} else if operation == "jmp" {
			i += value
		}
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
	instructions := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}

	for i, v := range instructions {
		op, _ := parseInstruction(v)
		if op == "jmp" {
			instructions[i] = "nop +0"
			success, accum := findInfiniteLoop(instructions)
			if success {
				fmt.Println(accum)
				break
			}
			instructions[i] = v
		}
	}
}
