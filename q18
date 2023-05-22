package main

import (
	"errors"
	"fmt"
)

func applyFunctionAndSum(slice []int, fn func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("o slice está vazio")
	}

	sum := 0
	for _, num := range slice {
		sum += fn(num)
	}

	return sum, nil
}

func double(n int) int {
	return n * 2
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	result, err := applyFunctionAndSum(numbers, double)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos resultados:", result)
	}
}
