package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/mowshon/iterium"
)

type IndexPair struct {
	y int
	x int
}

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
	width := len(data[0])
	visited := make(map[IndexPair]bool)
	positions := findAntennaPositions(data)

	// Get unique combinations of antenna pairs
	for _, indicesList := range positions {
		combinations := iterium.Combinations(indicesList, 2)
		toSlice, _ := combinations.Slice()

		for _, indicePair := range toSlice {
			p1 := indicePair[0]
			p2 := indicePair[1]

			xDiff := int(math.Abs(float64(p1.x - p2.x)))
			yDiff := int(math.Abs(float64(p1.y - p2.y)))

			epoch := int(math.Min(float64(width/xDiff), float64(height/yDiff)))
			for i := range epoch + 1 {
				// For each pair, validate potential antinode positions
				if p1.x < p2.x {
					if p1.y < p2.y {
						sum += checkValid(p1.x-i*xDiff, p1.y-i*yDiff, width, height, visited)
						sum += checkValid(p2.x+i*xDiff, p2.y+i*yDiff, width, height, visited)
					} else {
						sum += checkValid(p1.x-i*xDiff, p1.y+i*yDiff, width, height, visited)
						sum += checkValid(p2.x+i*xDiff, p2.y-i*yDiff, width, height, visited)
					}
				} else {
					if p1.y < p2.y {
						sum += checkValid(p1.x+i*xDiff, p1.y-i*yDiff, width, height, visited)
						sum += checkValid(p2.x-i*xDiff, p2.y+i*yDiff, width, height, visited)
					} else {
						sum += checkValid(p1.x+i*xDiff, p1.y+i*yDiff, width, height, visited)
						sum += checkValid(p2.x-i*xDiff, p2.y-i*yDiff, width, height, visited)
					}
				}
			}
		}
	}
	fmt.Println("Part Two:", sum)
}

func partOne(data []string) {
	sum := 0
	height := len(data)
	width := len(data[0])
	visited := make(map[IndexPair]bool)
	positions := findAntennaPositions(data)

	// Get unique combinations of antenna pairs
	for _, indicesList := range positions {
		combinations := iterium.Combinations(indicesList, 2)
		toSlice, _ := combinations.Slice()

		for _, indicePair := range toSlice {
			p1 := indicePair[0]
			p2 := indicePair[1]

			xDiff := int(math.Abs(float64(p1.x - p2.x)))
			yDiff := int(math.Abs(float64(p1.y - p2.y)))

			// For each pair, validate potential antinode positions
			if p1.x < p2.x {
				if p1.y < p2.y {
					sum += checkValid(p1.x-xDiff, p1.y-yDiff, width, height, visited)
					sum += checkValid(p2.x+xDiff, p2.y+yDiff, width, height, visited)
				} else {
					sum += checkValid(p1.x-xDiff, p1.y+yDiff, width, height, visited)
					sum += checkValid(p2.x+xDiff, p2.y-yDiff, width, height, visited)
				}
			} else {
				if p1.y < p2.y {
					sum += checkValid(p1.x+xDiff, p1.y-yDiff, width, height, visited)
					sum += checkValid(p2.x-xDiff, p2.y+yDiff, width, height, visited)
				} else {
					sum += checkValid(p1.x+xDiff, p1.y+yDiff, width, height, visited)
					sum += checkValid(p2.x-xDiff, p2.y-yDiff, width, height, visited)
				}
			}
		}
	}
	fmt.Println("Part One:", sum)
}

func checkValid(x int, y int, width int, height int, visited map[IndexPair]bool) int {
	// Skip visited antinodes
	if _, ok := visited[IndexPair{x, y}]; ok {
		return 0
	}

	if y < height && x < width && y >= 0 && x >= 0 {
		visited[IndexPair{x, y}] = true
		return 1
	}
	return 0
}

// Find antenna positions from the grid as a IndexPair (y, x)
func findAntennaPositions(data []string) map[string][]IndexPair {
	positions := make(map[string][]IndexPair)

	for i, row := range data {
		for j := range row {
			key := string(data[i][j])
			if key == "." {
				continue
			}
			if indicesList, ok := positions[key]; ok {
				positions[key] = append(indicesList, IndexPair{i, j})
			} else {
				positions[key] = []IndexPair{{i, j}}
			}
		}
	}
	return positions
}
