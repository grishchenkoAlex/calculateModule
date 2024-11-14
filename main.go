package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {
	operation := strings.Split(t.Content, "\"")

	if len(operation) != 5 && len(operation) != 3 {
		panic("Выражение не соответсвует требуемой арифметической операции")
	}

	if operation[0] != "" || len(operation) == 1 || operation[1] == "" {
		panic("При вводе строки должны быть в \"\" ")
	}

	signOperation := strings.ReplaceAll(operation[2], " ", "")
	firstString := operation[1]

	if len(operation) == 5 {
		secondString := operation[3]

		switch {
		case len(firstString) >= 10 || len(secondString) >= 10:
			panic("Максимальное количество символов строки не более 10")
		case signOperation == "+":
			fmt.Printf("%q\n", firstString+secondString)
		case signOperation == "-":
			difference := strings.Replace(firstString, secondString, "", 1)
			fmt.Printf("%q\n", difference)
		default:
			panic("Выражение не соответсвует требуемой арифметической операции")
		}
	}

	if len(operation) == 3 {
		var sign string
		var num int
		var resultString string
		var resultRune []rune
		firstStringRune := []rune(operation[1])

		if len(signOperation) == 2 || len(signOperation) == 3 {
			for i, val := range signOperation {
				if i == 0 {
					sign = string(val)
				} else {
					num, _ = strconv.Atoi(strings.Replace(signOperation, sign, "", 1))
				}
			}
		}

		switch {
		case num == 0 || num > 10:
			panic("число не соответствует требованиям")
		case sign == "/":
			for i := 0; i <= len(firstString)/num-1; i++ {
				resultRune = append(resultRune, firstStringRune[i])
			}
			fmt.Printf("%q\n", string(resultRune))
		case sign == "*":
			var i = 1
			for i <= num {
				resultString += firstString
				i++
			}
			if len(resultString) >= 40 {
				resultRune = []rune(resultString)
				resultRune = resultRune[0:40]
				fmt.Printf("%q\n", string(resultRune)+"...")
			} else {
				fmt.Printf("%q\n", resultString)
			}
		default:
			panic("Выражение не соответсвует требуемой арифметической операции")
		}
	}
}
func main() {
	text := &Text{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите выражение:")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}
