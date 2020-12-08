package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Luggage struct {
	bag         string
	description []BagDescription
}

type BagDescription struct {
	num uint64
	bag string
}

func Day7() (err error) {
	fmt.Println("Running day 7 challenge...")

	path := "days/day7.txt"
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	var luggage []Luggage
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		line = strings.TrimSuffix(line, ".")
		parts := strings.SplitN(line, " bags contain ", 2)
		bag := parts[0]

		var bags []BagDescription
		contains := strings.Split(parts[1], ", ")
		for _, c := range contains {
			if c == "no other bags" {
				bags = append(bags, BagDescription{num: 0, bag: ""})
				continue
			}

			c = strings.TrimSuffix(c, " bag")
			c = strings.TrimSuffix(c, " bags")
			parts := strings.SplitN(c, " ", 2)
			times, err := strconv.ParseUint(parts[0], 10, 64)
			if err != nil {
				return err // how to return this
			}
			bag := parts[1]

			bags = append(bags, BagDescription{num: times, bag: bag})
		}

		luggage = append(luggage, Luggage{bag: bag, description: bags})
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning %q: %w", path, err)
	}

	part1 := findContainingBags(luggage, "shiny gold")
	_, err = fmt.Printf("part1: %d\n", len(part1))
	if err != nil {
		return err
	}

	part2 := findRequiredBags(luggage, "shiny gold")
	_, err = fmt.Printf("part2: %d\n", part2)
	if err != nil {
		return err
	}

	return nil
}

// How many bag colors can eventually contain at least one shiny gold bag?
func findContainingBags(luggage []Luggage, needle string) map[string]bool {
	found := make(map[string]bool)
	for _, l := range luggage {
		if contains(l.description, needle) {
			bags := findContainingBags(luggage, l.bag)

			// add bags to found
			for k, v := range bags {
				found[k] = v
			}
			// add outer bag to found
			found[l.bag] = true
		}
	}

	return found
}

func contains(bags []BagDescription, needle string) bool {
	for _, b := range bags {
		if b.bag == needle {
			return true
		}
	}

	return false
}

// How many individual bags are required inside your single shiny gold bag?
func findRequiredBags(luggage []Luggage, bag string) uint64 {
	var total uint64

	descriptions := findBag(luggage, bag)
	if descriptions == nil {
		return total
	}

	for _, d := range descriptions {
		count := findRequiredBags(luggage, d.bag)
		count += 1            // include the found bag
		count = d.num * count // multiply the number of bags by the amount in the luggage

		total += count
	}

	return total
}

func findBag(luggage []Luggage, bag string) []BagDescription {
	for _, l := range luggage {
		if l.bag == bag {
			return l.description
		}
	}

	return []BagDescription{}
}

