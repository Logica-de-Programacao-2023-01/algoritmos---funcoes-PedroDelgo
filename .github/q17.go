package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func concatenateAndSort(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("o slice está vazio")
	}

	sort.Strings(slice)
	result := strings.Join(slice, "")

	return result, nil
}

func main() {
	stringsSlice := []string{"banana", "abacaxi", "laranja", "uva"}

	result, err := concatenateAndSort(stringsSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("String resultante:", result)
	}
}
