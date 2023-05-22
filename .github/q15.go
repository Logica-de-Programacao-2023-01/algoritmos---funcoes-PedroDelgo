package main

import (
	"errors"
	"fmt"
)

func containsNumber(number int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("o slice está vazio")
	}

	for _, value := range slice {
		if value == number {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	number := 3
	slice := []int{1, 2, 3, 4, 5}

	found, err := containsNumber(number, slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Número encontrado:", found)
	}
}
