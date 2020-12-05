package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

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

func validYear(y string, min int, max int) bool {
	if len(y) != 4 {
		return false
	}
	year, err := strconv.Atoi(y)
	if err != nil {
		return false
	}
	return year >= min && year <= max
}

func (p passport) ByrValid() bool {
	return validYear(p.Byr, 1920, 2002)
}

func (p passport) IyrValid() bool {
	return validYear(p.Iyr, 2010, 2020)
}

func (p passport) EyrValid() bool {
	return validYear(p.Eyr, 2020, 2030)
}

func (p passport) HgtValid() bool {
	hgtLen := len(p.Hgt)
	if hgtLen < 3 {
		return false
	}
	hgt, err := strconv.Atoi(p.Hgt[:hgtLen-2])
	if err != nil {
		return false
	}
	unit := p.Hgt[hgtLen-2:]
	if unit == "cm" {
		return hgt >= 150 && hgt <= 193
	}
	if unit == "in" {
		return hgt >= 59 && hgt <= 76
	}
	return false
}

func (p passport) HclValid() bool {
	if len(p.Hcl) != 7 {
		return false
	}
	if string(p.Hcl[0]) != "z" {
		return false
	}
	for i := 1; i < len(p.Hcl); i++ {
		found, err := regexp.MatchString("[a-z0-9]", string(p.Hcl[i]))
		if err != nil || !found {
			return false
		}
	}
	return true
}

func (p passport) EclValid() bool {
	validOptions := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	_, ok := validOptions[p.Ecl]
	return ok
}

func (p passport) PidValid() bool {
	if len(p.Pid) != 9 {
		return false
	}
	for i := 0; i < len(p.Pid); i++ {
		if !unicode.IsDigit(rune(p.Pid[i])) {
			return false
		}
	}
	return true
}

func (p passport) Valid() bool {
	return p.ByrValid() && p.IyrValid() && p.EyrValid() && p.HgtValid() && p.HclValid() && p.EclValid() && p.PidValid()
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
		y := strings.ReplaceAll(strings.ReplaceAll(s, ":", ": "), "#", "z")
		p := passport{}
		yaml.Unmarshal([]byte(y), &p)
		if p.Valid() {
			count++
		}
	}
	fmt.Println(count)
}
