package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type passport struct {
	Byr string `yaml:"byr"`
	Iyr string `yaml:"iyr"`
	Eyr string `yaml:"eyr"`
	Hgt string `yaml:"hgt"`
	Hcl string `yaml:"hcl"`
	Ecl string `yaml:"ecl"`
	Pid string `yaml:"pid"`
	Cid string `yaml:"cid"`
}

func (p passport) Valid() bool {
	return len(p.Byr) > 0 && len(p.Iyr) > 0 && len(p.Eyr) > 0 && len(p.Hgt) > 0 && len(p.Hcl) > 0 && len(p.Ecl) > 0 && len(p.Pid) > 0
}

func main() {
	var filePath string
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	} else {
		log.Fatal("Please provide a file path as an argument")
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	arr := strings.Split(str, "\n\n")
	whitespace := regexp.MustCompile(`\s+`)
	count := 0
	for _, val := range arr {
		s := whitespace.ReplaceAllString(val, "\n")
		y := strings.ReplaceAll(strings.ReplaceAll(s, ":", ": "), "#", "")
		p := passport{}
		yaml.Unmarshal([]byte(y), &p)
		if p.Valid() {
			count++
		}
	}
	fmt.Println(count)
}
