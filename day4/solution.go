package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func prepData() []string {
	var data []string

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, string(line))
	}
	return data
}

func main() {
	data := prepData()
	partOne(data)
	partTwo(data)
}

func partTwo(data []string) {
	sum := 0
	height := len(data)
	length := len(data[0])

	// Check everyposition and see if xmas or samx is present in possible positions.
	for i, row := range data {
		for j, _ := range row {
			// Check diagnoal right
			if j+3 <= length && i+2 <= height-1 {
				d1 := string(data[i][j]) + string(data[i+1][j+1]) + string(data[i+2][j+2])
				d2 := string(data[i][j+2]) + string(data[i+1][j+1]) + string(data[i+2][j])
				fmt.Println("d1", i, j, d1)
				fmt.Println("d2", i, j, d2)
				sum += checkSubstringPartTwo(d1, d2)
			}
		}
	}
	fmt.Println("SCORE", sum)
}

func checkSubstringPartTwo(str1 string, str2 string) int {
	if (str1 == "MAS" || str1 == "SAM") && (str2 == "MAS" || str2 == "SAM") {
		return 1
	}
	return 0
}

func partOne(data []string) {
	sum := 0
	height := len(data)
	length := len(data[0])

	// Check everyposition and see if xmas or samx is present in possible positions.
	for i, row := range data {
		for j, _ := range row {
			// Scan horizontally
			if j+4 <= length {
				sum += checkSubstring(row[j : j+4])
			}
			// Check vertical
			if i+3 <= height-1 {
				subStr := string(data[i][j]) + string(data[i+1][j]) + string(data[i+2][j]) + string(data[i+3][j])
				sum += checkSubstring(subStr)
			}
			// Check diagnoal right
			if j+4 <= length && i+3 <= height-1 {
				subStr := string(data[i][j]) + string(data[i+1][j+1]) + string(data[i+2][j+2]) + string(data[i+3][j+3])
				sum += checkSubstring(subStr)
			}
			// Check diagnoal left
			if j-3 >= 0 && i+3 <= height-1 {
				subStr := string(data[i][j]) + string(data[i+1][j-1]) + string(data[i+2][j-2]) + string(data[i+3][j-3])
				//fmt.Println("DIAG", i, j, subStr)
				sum += checkSubstring(subStr)
			}
		}
	}
	fmt.Println("SCORE", sum)
}

func checkSubstring(str string) int {
	if str == "XMAS" || str == "SAMX" {
		return 1
	}
	return 0
}
