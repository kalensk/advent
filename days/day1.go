package days

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)


func Day1() (err error) {
	fmt.Println("Running day 1 challenge...")

	path := "days/day1.txt"
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	var nums []uint64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			return err
		}

		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning %q: %w", path, err)
	}

	part1, err := sumOfTwoEntries(2020, nums)
	_, err = fmt.Printf("part1: %d\n", part1)
	if err != nil {
		return err
	}

	part2, err := sumOfThreeEntries(2020, nums)
	_, err = fmt.Printf("part2: %d\n", part2)
	if err != nil {
		return err
	}

	return nil
}

// Find the two entries that sum to 2020 and then multiply those two numbers
// together.
func sumOfTwoEntries(sum uint64, nums []uint64) (uint64, error) {
	for _, current := range nums {
		for _, next := range nums[1:] {
			if (current + next) == sum {
				return current * next, nil
			}
		}
	}

	return 0, errors.New("not found")
}

// Find the three entries that sum to 2020 and then multiply those three numbers
// together.
func sumOfThreeEntries(sum uint64, nums []uint64) (uint64, error) {
	for _, current := range nums {
		for _, next := range nums[1:] {
			for _, following := range nums[2:] {
				if (current + next + following) == sum {
					return current * next * following, nil
				}
			}
		}
	}

	return 0, errors.New("not found")
}
