package main

func part1(input string) int {
	grid := parseInput(input)
	node := parseInput(input)

	for row, line := range grid {
		for col, char := range line {
			if char != '.' {
				countAntinode(grid, node, char, row, col)
			}
		}
	}

	count := 0
	for _, r := range node {
		for _, c := range r {
			if c == '#' {
				count++
			}
		}
	}

	return count
}

func part2(input string) int {
	grid := parseInput(input)
	node := parseInput(input)

	for row, line := range grid {
		for col, char := range line {
			if char != '.' {
				markAntinodesResonantHarmonics(grid, node, char, row, col)
			}
		}
	}

	count := 0
	for _, r := range node {
		for _, c := range r {
			if c != '.' {
				count++
			}
		}
		println(string(r))
	}

	return count
}

func countAntinode(runes [][]rune, node [][]rune, r rune, si, sj int) {
	for row := si; row < len(runes); row++ {
		for col := 0; col < len(runes[row]); col++ {
			if runes[row][col] == r {
				diffI := row - si
				diffJ := col - sj

				if diffI == 0 && diffJ == 0 {
					continue
				}

				nextI := row + diffI
				nextJ := col + diffJ
				if inRange(node, nextI, nextJ) {
					node[nextI][nextJ] = '#'
				}

				prevI := si - diffI
				prevJ := sj - diffJ
				if inRange(node, prevI, prevJ) {
					node[prevI][prevJ] = '#'
				}
			}
		}
	}
}

func markAntinodesResonantHarmonics(grid [][]rune, node [][]rune, r rune, si, sj int) {
	for row := si; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == r {
				diffI := row - si
				diffJ := col - sj

				if diffI == 0 && diffJ == 0 {
					continue
				}

				nextI := row + diffI
				nextJ := col + diffJ
				for inRange(node, nextI, nextJ) {
					node[nextI][nextJ] = '#'
					nextI += diffI
					nextJ += diffJ
				}

				prevI := si - diffI
				prevJ := sj - diffJ
				for inRange(node, prevI, prevJ) {
					node[prevI][prevJ] = '#'
					prevI -= diffI
					prevJ -= diffJ
				}
			}
		}
	}
}

func inRange(runes [][]rune, i, j int) bool {
	return 0 <= i && i < len(runes) && 0 <= j && j < len(runes[i])
}

func parseInput(input string) [][]rune {
	grid := make([][]rune, 0)
	line := make([]rune, 0)

	for _, c := range input {
		switch c {
		case '\n':
			grid = append(grid, line)
			line = make([]rune, 0)
		default:
			line = append(line, c)
		}
	}

	if len(line) > 0 {
		grid = append(grid, line)
	}
	return grid
}
