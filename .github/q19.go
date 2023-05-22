package main

import (
	"errors"
	"fmt"
)

func findPrimes(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("o número é negativo")
	}

	primes := make([]int, 0)
	isPrime := true

	for i := 2; i <= n; i++ {
		isPrime = true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes, nil
}

func main() {
	number := 20

	result, err := findPrimes(number)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos:", result)
	}
}
