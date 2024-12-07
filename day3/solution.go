package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	var sum int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
		matches := r.FindAllStringSubmatch(line, -1)

		for _, v := range matches {
			n1, _ := strconv.Atoi(v[1])
			n2, _ := strconv.Atoi(v[2])
			sum += n1 * n2
		}

	}
	fmt.Println(sum)
}

func partTwo() {
	var sum int
	var inputStr string

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputStr += line
	}

	// Match any mul(x, y) pattern between `Don't()` and `Do()` and remove them from the string
	ignorePatterm := `don\'t\(\).+?.+?do\(\)`
	ignoreRe := regexp.MustCompile(ignorePatterm)
	matchesToIgnore := ignoreRe.FindAllStringSubmatch(inputStr, -1)

	for _, v := range matchesToIgnore {
		inputStr = strings.Replace(inputStr, v[0], "", -1)
	}

	r := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	matches := r.FindAllStringSubmatch(inputStr, -1)

	for _, v := range matches {
		n1, _ := strconv.Atoi(v[1])
		n2, _ := strconv.Atoi(v[2])
		sum += n1 * n2
	}
	fmt.Println(sum)
}
