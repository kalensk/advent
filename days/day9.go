package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day9() (err error) {
	fmt.Println("Running day 9 challenge...")

	path := "days/day9.txt"
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
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		num, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			return err
		}

		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning %q: %w", path, err)
	}

	part1, err := XMASNotSumOfPreviousNumbers(nums, 25)
	_, err = fmt.Printf("part1: %d\n", part1)
	if err != nil {
		return err
	}

	part2, err := findEncryptionWeakness(nums, part1)
	_, err = fmt.Printf("part2: %d\n", part2)
	if err != nil {
		return err
	}

	return nil
}

// Find the first number in the list (after the preamble) which is not the sum
// of two of the 25 numbers before it. What is the first number that does not
// have this property?
func XMASNotSumOfPreviousNumbers(nums []uint64, offset uint64) (uint64, error) {
	for int(offset) < len(nums) {
		found := findNumsAddingToTarget(nums[:offset], nums[offset])
		if !found {
			return nums[offset], nil
		}

		offset++
	}

	return 0, nil
}

// given a list of numbers, which 2 unique numbers add to the target number
func findNumsAddingToTarget(nums []uint64, target uint64) bool {
	for _, num1 := range nums {
		for _, num2 := range nums {
			if (num1 + num2) == target {
				return true
			}
		}
	}

	return false
}

// What is the encryption weakness in your XMAS-encrypted list of numbers?
func findEncryptionWeakness(nums []uint64, target uint64) (uint64, error) {
	var offset uint64
	var subOffset uint64
	var sum uint64

	for int(offset) < len(nums) {
		for sum < target {
			sum += nums[subOffset]
			if sum == target {
				subRange := nums[offset:subOffset]
				return sumSmallestAndLargestNumInRange(subRange), nil
			}
			subOffset++
		}

		sum = 0
		offset++
		subOffset = offset
	}

	return 0, nil
}

// To find the encryption weakness, add together the smallest and largest number
// in this contiguous range; in this example, these are 15 and 47, producing 62.
func sumSmallestAndLargestNumInRange(nums []uint64) uint64 {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return nums[0] + nums[len(nums)-1]
}
