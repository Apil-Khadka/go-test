package main

import "fmt"

func main() {
	// Set up a slice to store the prime numbers
	var primes []int

	// Start with the number 2, which is the first prime number
	num := 2

	// Keep looping until we have found 100 prime numbers
	for len(primes) < 10 {
		// Assume that the number is prime
		isPrime := true

		// Check if the number is divisible by any of the prime numbers we have found so far
		for _, prime := range primes {
			if num%prime == 0 {
				// If it is divisible, it is not a prime number
				isPrime = false
				break
			}
		}

		// If the number is prime, add it to the slice
		if isPrime {
			primes = append(primes, num)
		}

		// Move on to the next number
		num++
	}

	// Print the prime numbers
	for _, prime := range primes {
		fmt.Println(prime)
	}
}
