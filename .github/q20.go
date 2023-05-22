package main

import (
	"errors"
	"fmt"
)

func filterStrings(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("o slice está vazio")
	}

	filtered := make([]string, 0)

	for _, str := range slice {
		if len(str) > 5 {
			filtered = append(filtered, str)
		}
	}

	return filtered, nil
}

func main() {
	stringsSlice := []string{"apple", "banana", "orange", "grapefruit", "kiwi"}

	result, err := filterStrings(stringsSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings filtradas:", result)
	}
}
