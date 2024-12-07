package main

import (
	"strconv"
)

func part1(input string) int {
	eqs := make([][]int, 0)

	eq := make([]int, 0)
	s := make([]rune, 0)
	for _, c := range input {
		if c == '\n' {
			if len(s) > 0 {
				ls, _ := strconv.Atoi(string(s))
				eq = append(eq, ls)
			}
			eqs = append(eqs, eq)
			eq = make([]int, 0)
			s = make([]rune, 0)
		} else if c == ' ' {
			if len(s) > 0 {
				ls, _ := strconv.Atoi(string(s))
				eq = append(eq, ls)
				s = make([]rune, 0)
			}
		} else if c == ':' {
			if len(s) > 0 {
				ls, _ := strconv.Atoi(string(s))
				eq = append(eq, ls)
				s = make([]rune, 0)
			}
		} else {
			s = append(s, c)
		}
	}
	if len(eq) > 0 {
		if len(s) > 0 {
			ls, _ := strconv.Atoi(string(s))
			eq = append(eq, ls)
		}
		eqs = append(eqs, eq)
	}

	valid := make([]int, 0)
	for _, eq := range eqs {
		if len(eq) < 3 {
			continue
		}

		sample := 1 << (len(eq) - 2) // 2^(n-2) combinations
		for i := 0; i < sample; i++ {
			result := eq[1]
			for j := 0; j < len(eq)-2; j++ {
				if result > eq[0] {
					break
				}

				if (i>>j)&1 == 0 {
					result = result + eq[j+2]
				} else {
					result = result * eq[j+2]
				}
			}

			if result == eq[0] {
				valid = append(valid, eq[0])
				break
			}
		}
	}

	sum := 0
	for _, v := range valid {
		sum = add(sum, v)
	}

	return sum
}

func part2(input string) int {
	return 0
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}
