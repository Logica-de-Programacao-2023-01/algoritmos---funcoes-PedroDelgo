package main

import (
	"errors"
	"fmt"
	"strings"
)

func replaceVowels(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("a string está vazia")
	}

	vowels := "aeiouAEIOU"
	result := strings.Builder{}

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			result.WriteRune('*')
		} else {
			result.WriteRune(char)
		}
	}

	return result.String(), nil
}

func main() {
	input := "Hello, World!"

	replaced, err := replaceVowels(input)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("String substituída:", replaced)
	}
}
