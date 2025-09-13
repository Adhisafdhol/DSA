package main

import "fmt"

// Find Gthe greatest common divisor
func GCD(p int, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return GCD(q, r)
}

func main() {
	result := GCD(123456789, 789)

	fmt.Println(result)
}
