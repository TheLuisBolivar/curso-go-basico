package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()
	valores := strings.Split(operation, "+")
	var number1 int = validateValue(valores[0])
	var number2 int = validateValue(valores[1])
	var result = number1 + number2
	fmt.Println(result)
}

func validateValue(value string) int {
	number, err := strconv.Atoi(value)
	var resutl int = 0

	if err != nil {
		fmt.Println("Se encontro valor en los valores ingresados {}", err)
	} else {
		resutl = number
	}

	return resutl
}
