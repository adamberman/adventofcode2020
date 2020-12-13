package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ruleToKeyAndCount(rule string) (string, int) {
	ruleWords := strings.Split(rule, " ")
	bag := []string{ruleWords[1], ruleWords[2]}
	count, err := strconv.Atoi(ruleWords[0])
	if err != nil {
		log.Fatal(err)
	}
	return strings.Join(bag, " "), count
}

func processRule(s string) (string, []string) {
	arr := strings.Split(s, " bags contain ")
	bagType := arr[0]
	rawRule := arr[1]
	rule := []string{}
	if rawRule != "no other bags." {
		rule = strings.Split(rawRule, ", ")
	}

	return bagType, rule
}

func dfsBagCount(startingBag string, rules map[string][]string) int {
	var totalBags int
	nextBagsToCheck, ok := rules[startingBag]
	if ok {
		for _, rule := range nextBagsToCheck {
			nextBag, count := ruleToKeyAndCount(rule)
			fmt.Println(nextBag)
			fmt.Println(count)
			totalBags += count
			totalBags += count * dfsBagCount(nextBag, rules)
		}
	}

	return totalBags
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

	bagTypes := make(map[string]bool)
	rules := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bagType, rule := processRule(line)
		_, exists := bagTypes[bagType]
		if !exists {
			bagTypes[bagType] = true
			rules[bagType] = rule
		}
	}

	startingBag := "shiny gold"
	count := dfsBagCount(startingBag, rules)
	fmt.Println(count)
}
