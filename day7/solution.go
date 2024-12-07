package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int
	var sum2 int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(res[0])

		numSplit := strings.Split(res[1], " ")
		var arr = []int{}

		for _, i := range numSplit {
			n, _ := strconv.Atoi(i)
			arr = append(arr, n)
		}
		// Part one
		solvable := solve(arr, testValue)
		if solvable {
			sum += testValue
		}

		// Part two
		solvable2 := solvePartTwo(arr, testValue)
		if solvable2 {
			sum2 += testValue
		}

	}
	fmt.Println("Part One: ", sum)
	fmt.Println("Part Two: ", sum2)
}

func solve(arr []int, testValue int) bool {
	length := len(arr)

	if length == 2 {
		if (arr[0]*arr[1] == testValue) || (arr[0]+arr[1] == testValue) {
			return true
		}
		return false
	}

	// Reduce array with first two numbers multiplied
	arr1 := make([]int, 0, length-1)
	arr1 = append(arr1, arr[0]*arr[1])
	arr1 = append(arr1, arr[2:]...)

	// Reduce array with first two numbers added
	arr2 := make([]int, 0, length-1)
	arr2 = append(arr2, arr[0]+arr[1])
	arr2 = append(arr2, arr[2:]...)

	s1 := solve(arr1, testValue)
	s2 := solve(arr2, testValue)
	return s1 || s2
}

func solvePartTwo(arr []int, testValue int) bool {
	length := len(arr)

	str := fmt.Sprintf("%d%d", arr[0], arr[1])
	concatNum, _ := strconv.Atoi(str)

	if length == 2 {
		if (arr[0]*arr[1] == testValue) || (arr[0]+arr[1] == testValue) || (concatNum == testValue) {
			return true
		}
		return false
	}

	// Reduce array with first two numbers multiplied
	arr1 := make([]int, 0, length-1)
	arr1 = append(arr1, arr[0]*arr[1])
	arr1 = append(arr1, arr[2:]...)

	// Reduce array with first two numbers added
	arr2 := make([]int, 0, length-1)
	arr2 = append(arr2, arr[0]+arr[1])
	arr2 = append(arr2, arr[2:]...)

	// Reduce array with first two numbers concatenated
	arr3 := make([]int, 0, length-1)
	arr3 = append(arr3, concatNum)
	arr3 = append(arr3, arr[2:]...)

	s1 := solvePartTwo(arr1, testValue)
	s2 := solvePartTwo(arr2, testValue)
	s3 := solvePartTwo(arr3, testValue)
	return s1 || s2 || s3
}
