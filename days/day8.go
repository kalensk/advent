package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op      string
	arg     int64
	visited bool
}

func Day8() (err error) {
	fmt.Println("Running day 8 challenge...")

	path := "days/day8.txt"
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	var instructions []*instruction
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		op := parts[0]
		arg, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return err
		}

		instructions = append(instructions, &instruction{op: op, arg: arg})
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning %q: %w", path, err)
	}

	part1, err := returnAccBeforeInfiniteLoop(instructions)
	_, err = fmt.Printf("part1: %d\n", part1)
	if err != nil {
		return err
	}

	resetProgram(instructions, -1)

	part2, err := fixProgram(instructions)
	_, err = fmt.Printf("part2: %d\n", part2)
	if err != nil {
		return err
	}

	resetProgram(instructions, -1)

	return nil
}

// Immediately before any instruction is executed a second time, what value is
// in the accumulator?
func returnAccBeforeInfiniteLoop(instructions []*instruction) (int64, error) {
	var acc int64
	var pos int64

	for pos <= int64(len(instructions)) {
		i := instructions[pos]

		if i.visited == true {
			return acc, nil
		}

		switch i.op {
		case "acc":
			acc += i.arg
			pos++
			i.visited = true
		case "jmp":
			pos += i.arg
			i.visited = true
		case "nop":
			pos++
			i.visited = true
		}
	}

	return acc, nil
}

func returnOrder(instructions []*instruction) []instruction {
	return nil
}

// Fix the program so that it terminates normally by changing exactly one jmp
// (to nop) or nop (to jmp). What is the value of the accumulator after the
// program terminates?
func fixProgram(instructions []*instruction) (int64, error) {
	for pos, i := range instructions {
		switch i.op {
		case "jmp":
			i.op = "nop"
		case "nop":
			i.op = "jmp"
		default:
			continue
		}

		ran, acc := runProgram(instructions)
		if ran {
			return acc, nil
		}

		resetProgram(instructions, pos)
	}

	return 0, nil
}

func runProgram(instructions []*instruction) (bool, int64) {
	var acc int64
	var pos int64

	for pos < int64(len(instructions)) {
		i := instructions[pos]

		if i.visited == true {
			return false, 0
		}

		switch i.op {
		case "acc":
			acc += i.arg
			pos++
			i.visited = true
		case "jmp":
			pos += i.arg
			i.visited = true
		case "nop":
			pos++
			i.visited = true
		}
	}

	return true, acc
}

func resetProgram(instructions []*instruction, pos int) {
	for _, i := range instructions {
		i.visited = false
	}

	if pos == -1 {
		return
	}

	i := instructions[pos]
	switch i.op {
	case "jmp":
		i.op = "nop"
	case "nop":
		i.op = "jmp"
	}
}
