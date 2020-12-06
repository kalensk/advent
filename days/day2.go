package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	policy Policy
	pw     string
}

// Password policy is how many min and max times the given letter must appear
// in a password to be valid.
type Policy struct {
	letter string
	lower  int
	upper  int
}

func Day2() (err error) {
	fmt.Println("Running day 2 challenge...")

	path := "days/day2.txt"
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	var pws []PasswordPolicy
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(line, ": ", 2)
		policy := parts[0]
		pw := parts[1]

		policyParts := strings.Fields(policy)
		letter := policyParts[1]
		times := strings.SplitN(policyParts[0], "-", 2)
		lower, err := strconv.Atoi(times[0])
		if err != nil {
			return err
		}

		upper, err := strconv.Atoi(times[1])
		if err != nil {
			return err
		}

		pwPolicy := PasswordPolicy{
			policy: Policy{
				letter: letter,
				lower:  lower,
				upper:  upper,
			},
			pw:     pw,
		}

		pws = append(pws, pwPolicy)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning %q: %w", path, err)
	}

	part1, err := validPasswordsPart1(pws)
	_, err = fmt.Printf("part1: %d\n", part1)
	if err != nil {
		return err
	}

	part2, err := validPasswordsPart2(pws)
	_, err = fmt.Printf("part2: %d\n", part2)
	if err != nil {
		return err
	}

	return nil
}

// Given a password policy and password return how many valid passwords are in
// the input.
// The password policy indicates the lowest and highest number of times a given
// letter must appear for the password to be valid. For example, 1-3 a means that
// the password must contain a at least 1 time and at most 3 times.
func validPasswordsPart1(pws []PasswordPolicy) (int, error) {
	valid := 0
	for _, pw := range pws {
		count := strings.Count(pw.pw, pw.policy.letter)
		if count >= pw.policy.lower && count <= pw.policy.upper {
			valid++
		}
	}

	return valid, nil
}


// Given a password policy and password return how many valid passwords are in
// the input.
// Each policy actually describes two positions in the password, where 1 means
// the first character, 2 means the second character, and so on. (Be careful;
// Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of
// these positions must contain the given letter. Other occurrences of the letter
// are irrelevant for the purposes of policy enforcement.
func validPasswordsPart2(pws []PasswordPolicy) (int, error) {
	valid := 0
	for _, pw := range pws {
		letter := []rune(pw.policy.letter)[0]
		// passwords are not 0-based index
 		inPos1 := ([]rune(pw.pw))[pw.policy.lower - 1] == letter
		inPos2 := ([]rune(pw.pw))[pw.policy.upper - 1] == letter

		if (inPos1 && !inPos2) || (!inPos1 && inPos2) {
			valid++
		}
	}

	return valid, nil
}
