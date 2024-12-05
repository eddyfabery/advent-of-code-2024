package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var list1 []int
	var list2 []int
	var list2Counter map[int]int = map[int]int{}

	var n int
	var diffSum int
	var similarityScore int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, "   ")
		s1, _ := strconv.Atoi(res[0])
		s2, _ := strconv.Atoi(res[1])
		list1 = append(list1, s1)
		list2 = append(list2, s2)

		val, ok := list2Counter[s2]
		if ok {
			list2Counter[s2] = val + 1
		} else {
			list2Counter[s2] = 1
		}
		n += 1
	}
	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < n; i++ {
		diffSum += int(math.Abs(float64(list1[i] - list2[i])))

		val, ok := list2Counter[list1[i]]
		if ok {
			similarityScore += val * list1[i]
		}
	}
	fmt.Println(diffSum)
	fmt.Println(similarityScore)

}
