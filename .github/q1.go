package main

import "fmt"

//Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.

func func1(slice []int) float64 {
	tamanho := len(slice)
	soma := 0
	for _, valor := range slice {
		soma += valor
	}
	media := soma / tamanho
	return float64(media)
}
func main() {
	numeros := []int{5, 10, 15, 20, 25}
	media := func1(numeros)
	fmt.Printf("A média é: %.2f\n", media)
}
