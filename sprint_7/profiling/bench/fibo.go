package main

func FiboRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FiboRecursive(n-2) + FiboRecursive(n-1)
}
