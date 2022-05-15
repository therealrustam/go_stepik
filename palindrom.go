// Программа по определению слов палиндромов.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var result, word string
	fmt.Println("Введите слово:")
	fmt.Scan(&word)
	for _, v := range word {
		result = string(v) + result
	}
	if strings.ToLower(result) == strings.ToLower(word) {
		fmt.Print("Слово " + strings.ToLower(word) + " явлется палиндромом.")
	} else {
		fmt.Print("Слово " + strings.ToLower(word) + " не явлется палиндромом.")
	}
}
