package main

func part1(input string) int {
	grid, guard := parseInput(input)

	for y, x := guard.walk(); inRange(grid, y, x); y, x = guard.walk() {
		if grid[y][x] == '#' {
			guard.turn()
		} else {
			grid[guard.y][guard.x] = 'X'
			guard.y, guard.x = y, x
		}
	}

	grid[guard.y][guard.x] = 'X'

	count := 0
	for _, row := range grid {
		for _, r := range row {
			if r == 'X' {
				count++
			}
		}
	}

	return count
}

func part2(input string) int {
	grid, guard := parseInput(input)

	starterPoint := Guard{guard.y, guard.x, guard.dir}
	obstruction := make(map[[2]int]int)
	for y, x := guard.walk(); inRange(grid, y, x); y, x = guard.walk() {
		if grid[y][x] == '#' {
			guard.turn()
		} else {
			guard.y, guard.x = y, x
			if hasLoop(grid, Guard{starterPoint.y, starterPoint.x, starterPoint.dir}, [2]int{y, x}) {
				obstruction[[2]int{y, x}] = 1
			}
		}

	}
	return len(obstruction)
}

type Guard struct {
	y, x, dir int
}

var directions = [4][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func (g *Guard) turn() {
	g.dir = (g.dir + 1) % 4
}

func (g *Guard) walk() (int, int) {
	row := g.y + directions[g.dir][0]
	col := g.x + directions[g.dir][1]
	return row, col
}

func inRange(grid [][]rune, row, col int) bool {
	return 0 <= row && row < len(grid) && 0 <= col && col < len(grid[0])
}

func parseInput(input string) ([][]rune, Guard) {
	var grid [][]rune
	var guard Guard

	var runes []rune
	for _, r := range input {
		if r == '\n' {
			grid = append(grid, runes)
			runes = []rune{}
		} else {
			runes = append(runes, r)
			if r == '^' && guard.y == 0 {
				guard = Guard{len(grid), len(runes) - 1, 0}
			}
		}
	}
	if len(runes) > 0 {
		grid = append(grid, runes)
	}
	return grid, guard
}

func hasLoop(grid [][]rune, guard Guard, o [2]int) bool {
	visited := make(map[[3]int]bool)
	for y, x := guard.walk(); inRange(grid, y, x); y, x = guard.walk() {
		if grid[y][x] == '#' || (y == o[0] && x == o[1]) {
			guard.turn()
		} else {
			if visited[[3]int{y, x, guard.dir}] {
				return true
			}
			visited[[3]int{y, x, guard.dir}] = true
			guard.y, guard.x = y, x
		}
	}
	return false
}
