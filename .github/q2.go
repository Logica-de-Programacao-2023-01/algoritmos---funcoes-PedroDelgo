package main

func contarVogais(texto string) int {
	vogais := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	quantidade := 0

	for _, char := range texto {
		for _, vogal := range vogais {
			if char == vogal {
				quantidade++
				break
			}
		}
	}

	return quantidade
}
