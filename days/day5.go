package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day5() (err error) {
	fmt.Println("Running day 5 challenge...")

	path := "days/day5.txt"
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	var seats []int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		// convert to binary
		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "B", "1")

		line = strings.ReplaceAll(line, "R", "1")
		line = strings.ReplaceAll(line, "L", "0")

		num, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			return err
		}

		seats = append(seats, num)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning %q: %w", path, err)
	}

	part1, err := findHighestSeatID(seats)
	_, err = fmt.Printf("part1: %d\n", part1)
	if err != nil {
		return err
	}

	part2, err := seatID(seats)
	_, err = fmt.Printf("part2: %d\n", part2)
	if err != nil {
		return err
	}

	return nil
}

// What is the highest seat ID on a boarding pass?
func findHighestSeatID(seats []int64) (int64, error) {
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})

	return seats[len(seats)-1], nil
}

type Seat struct {
	row int64
	col int64
}

// Your seat wasn't at the very front or back, though; the seats with IDs +1
// and -1 from yours will be in your list.
// What is the ID of your seat?
func seatID(seats []int64) (int64, error) {
	// create a grid of filled seats
	filled := make(map[Seat]bool)
	for row := 0; row < 128; row++ {
		for col := 0; col < 8; col++ {
			if row == 0 || row == 127 {
				continue
			}
			filled[Seat{row: int64(row), col: int64(col)}] = true
		}
	}

	for _, seat := range seats {
		row := seat / 8
		col := seat % 8

		if row == 0 || row == 127 {
			continue
		}

		filled[Seat{row: row, col: col}] = false
	}

	for seat, filled := range filled {
		if filled == false {
			// to fix
			return (seat.row * 8) + seat.col / 8, nil
		}
	}

	return 0, nil
}
