package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution1 := firstProblem()
	fmt.Println("Solution1:", solution1)

	solution2 := secondProblem()
	fmt.Println("Solution2:", solution2)
}

func firstProblem() int {
	var safeCount int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, " ")

		if isSafe(res) {
			safeCount += 1
		}
	}
	return safeCount
}

func secondProblem() int {
	var safeCount int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, " ")
		length := len(res)

		if isSafe(res) {
			safeCount += 1
		} else {
			for i := 0; i < length; i++ {
				newArr := make([]string, 0, length-1)
				newArr = append(newArr, res[:i]...)
				newArr = append(newArr, res[i+1:]...)

				if isSafe(newArr) {
					safeCount += 1
					break
				}
			}
		}
	}
	return safeCount
}

func isSafe(arr []string) bool {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		s1, _ := strconv.Atoi(arr[i])
		s2, _ := strconv.Atoi(arr[i+1])

		// check current value and next value
		diff := int(math.Abs(float64(s1 - s2)))
		if !(0 < diff && diff < 4) {
			return false
		}

		// check previous value with current value if previous exists
		if i-1 >= 0 {
			s0, _ := strconv.Atoi(arr[i-1])

			if !((s0 < s1 && s1 < s2) || (s0 > s1 && s1 > s2)) {
				return false
			}
			diff := int(math.Abs(float64(s0 - s1)))
			if !(0 < diff && diff < 4) {
				return false
			}
		}
	}

	return true
}
