// Программа по проверке паролей.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var first string
	fmt.Println("Введите пароль из комбинации цифр и букв латиского алфавита:")
	fmt.Scan(&first)
	counter := 0
	length := utf8.RuneCountInString(first)
	for _, v := range first {
		number := unicode.IsDigit(v)
		latin := unicode.Is(unicode.Latin, v)
		if number || latin {
			counter++
		}
	}
	if (counter == length) && (length >= 5) {
		fmt.Print("Пароль принят")
	} else {
		fmt.Print("Пароль не принят")
	}
}
