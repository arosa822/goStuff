package main

import "fmt"

// Find primes in the range 2..n and return them in a slice

func sieveOfEratosthenes(n int) []int {
	// Create a boolean array [0..n] as all true, and then
	// set entries false as they are found not to be prime

	integers := make([]bool, n+1) // why "n+1" and not "n"?

	for i := range integers {
		integers[i] = true
	}

	integers[0], integers[1] = false, false

	// we "cross out" multiples of each prime because they
	// are divisible by that prime number; e.g., when p =
	// 2 we remove 4, 6, 8 ... and when p = 3 we remove
	// 9,12,15,... and skip 6 because 6 = 2*3 was seen

	for p := 2; p*p <= n; p++ {
		// If integers[p] is not changed, then it is prime

		if integers[p] {
			// update values x, x+p, x+2p,.. as not prime
			// starting with x = p*p (see above)

			for i := p * p; i <= n; i += p {
				integers[i] = false
			}

		}
	}
	// Now pick out the primes and return only them
	var primes []int // we don't know how many yet

	for p := range integers {
		if integers[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	p := sieveOfEratosthenes(121)

	fmt.Printf("%d:%v\n", len(p), p)
}
