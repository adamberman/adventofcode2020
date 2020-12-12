package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func processRule(s string) (string, []string) {
	arr := strings.Split(s, " bags contain ")
	bagType := arr[0]
	rawRule := arr[1]
	rule := []string{}
	if rawRule != "no other bags." {
		subRules := strings.Split(rawRule, ", ")
		for _, r := range subRules {
			subRuleWords := strings.Split(r, " ")
			bagRuleWords := []string{subRuleWords[1], subRuleWords[2]}
			bagRule := strings.Join(bagRuleWords, " ")
			rule = append(rule, bagRule)
		}
	}

	return bagType, rule
}

func bfsBagSearch(start string, target string, rules map[string][]string) bool {
	queue := []string{start}
	seen := make(map[string]bool)
	seen[start] = true
	for i := 0; i < len(queue); i++ {
		key := queue[i]
		vals, ok := rules[key]
		if !ok {
			continue
		}
		for _, val := range vals {
			if val == target {
				return true
			}
			_, exists := seen[val]
			if !exists {
				queue = append(queue, val)
				seen[val] = true
			}
		}
	}
	return false
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
	var count int
	for k := range bagTypes {
		if bfsBagSearch(k, startingBag, rules) {
			count++
		}
	}
	fmt.Println(count)
}
