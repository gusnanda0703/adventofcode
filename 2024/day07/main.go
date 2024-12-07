package main

import (
	"strconv"
)

func part1(input string) int {
	eqs := parseInput(input)

	valid := make([]int, 0)
	for _, eq := range eqs {
		if len(eq) < 3 {
			continue
		}

		target, initial := eq[0], eq[1]
		operations := 1 << (len(eq) - 2) // 2^(n-2) combinations

		for i := 0; i < operations; i++ {
			result := initial
			for j := 0; j < len(eq)-2; j++ {
				if result > target {
					break
				}
				if (i>>j)&1 == 0 {
					result += eq[j+2]
				} else {
					result *= eq[j+2]
				}
			}
			if result == target {
				valid = append(valid, target)
				break
			}
		}
	}

	sum := 0
	for _, v := range valid {
		sum += v
	}

	return sum
}

func part2(input string) int {
	eqs := parseInput(input)

	valid := make([]int, 0)
	for _, eq := range eqs {
		if len(eq) < 3 {
			continue
		}

		target, initial := eq[0], eq[1]
		operations := pow(3, len(eq)-2)

		for i := 0; i < operations; i++ {
			result := initial

			num := i
			for j := 0; j < len(eq)-2; j++ {
				if result > target {
					break
				}

				op := num % 3
				switch op {
				case 0:
					result += eq[j+2]
				case 1:
					result *= eq[j+2]
				case 2:
					result = concatenateIntegers(result, eq[j+2])
				}

				num /= 3
			}
			if result == target {
				valid = append(valid, target)
				break
			}
		}
	}

	sum := 0
	for _, v := range valid {
		sum += v
	}

	return sum
}

func parseInput(input string) [][]int {
	eqs := make([][]int, 0)
	eq := make([]int, 0)
	s := make([]rune, 0)

	appendNumber := func() {
		if len(s) > 0 {
			ls, _ := strconv.Atoi(string(s))
			eq = append(eq, ls)
			s = make([]rune, 0)
		}
	}

	for _, c := range input {
		switch c {
		case '\n':
			appendNumber()
			eqs = append(eqs, eq)
			eq = make([]int, 0)
		case ' ', ':':
			appendNumber()
		default:
			s = append(s, c)
		}
	}
	appendNumber()
	if len(eq) > 0 {
		eqs = append(eqs, eq)
	}
	return eqs
}

func countDigits(num int) int {
	digits := 0
	for num > 0 {
		num /= 10
		digits++
	}
	// Handle case for num = 0
	if digits == 0 {
		return 1
	}
	return digits
}

func concatenateIntegers(nums ...int) int {
	result := 0
	for _, num := range nums {
		digits := countDigits(num)
		result *= pow(10, digits)
		result += num
	}
	return result
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
