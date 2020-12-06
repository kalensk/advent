package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day3() (err error) {
	fmt.Println("Running day 3 challenge...")

	path := "days/day3.txt"
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning %q: %w", path, err)
	}

	part1, err := findTreesPart1(lines, 3, 1)
	_, err = fmt.Printf("part1: %d\n", part1)
	if err != nil {
		return err
	}

	part2, err := findTreesPart2(lines)
	_, err = fmt.Printf("part2: %d\n", part2)
	if err != nil {
		return err
	}

	return nil
}

// Start by counting all the trees you would encounter for the slope right 3,
// down 1
func findTreesPart1(lines []string, xdelta int, ydelta int) (uint64, error) {
	xpos := 0
	ypos := 0
	var numTrees uint64

	for ypos < len(lines) {
		row := lines[ypos]
		elem := string(row[xpos%len(row)])

		if elem == "#" {
			numTrees++
		}

		xpos += xdelta
		ypos += ydelta
	}

	return numTrees, nil
}

/*
Determine the number of trees you would encounter if, for each of the following slopes, you start at the top-left corner and traverse the map all the way to the bottom:

Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.

What do you get if you multiply together the number of trees encountered on each of the listed slopes?
*/
func findTreesPart2(lines []string) (uint64, error) {
	type delta struct {
		x int
		y int
	}

	deltas := []delta{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}

	var total uint64 = 1
	for _, delta := range deltas {
		numTrees, err := findTreesPart1(lines, delta.x, delta.y)
		if err != nil {
			return 0, err
		}

		total *= numTrees
	}

	return total, nil
}
