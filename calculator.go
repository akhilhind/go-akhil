package main

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
