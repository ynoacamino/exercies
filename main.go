package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(sumParNums(10, 50))
	fmt.Println(invertString("Hola mundo"))

	response, err := factorial(5)

	if err != nil {
		panic(err)
	}

	fmt.Println("Ejercicio 3: factorial de un numero")
	fmt.Println(response)

	fmt.Println(isPrimo(103))

	fmt.Println(findIndex([]int{0, 1, 2, 3, 4, 5}, 3))

	fmt.Println(isParindromo("holaloh"))
	fmt.Println(isParindromo("hola"))

	fmt.Println(random(100))
	fmt.Println(random(5))

	chapter()
}

// no autocompletes copilot
func sumParNums(a int, b int) int {
	println("Ejercicio 1: suma de numeros pares dentro de un rango")

	var sum int

	for i := a; i <= b; i += 1 {
		if i%2 == 0 {
			sum += i
		}
	}

	return sum
}

func invertString(s string) string {
	println("Ejercicio 2: invertir un string")
	var inverted string

	for i := len(s) - 1; i >= 0; i -= 1 {
		inverted += string(s[i])
	}

	return inverted
}

func factorial(num int32) (response int32, err error) {
	if num < 0 {
		return 0, errors.New("negative number")
	}

	if num == 0 {
		return 1, nil
	}
	if num == 1 {
		return 1, nil
	}

	fact, err := factorial(num - 1)

	if err != nil {
		panic(err)
	}

	return fact * num, nil
}

func isPrimo(num int) bool {
	fmt.Println("Ejercicio 4: numero primo")

	for i := 2; i < num; i += 1 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findIndex(arr []int, element int) int {
	println("Ejercicio 5: buscar un elemento")

	for i := 0; i < len(arr); i += 1 {
		if arr[i] == element {
			return i
		}
	}

	return -1
}

func isParindromo(s string) bool {
	println("Ejercicio 6: es palindromo")
	var length int

	if len(s)%2 == 0 {
		length = len(s) / 2
	} else {
		length = (len(s) - 1) / 2
	}

	for i := 0; i < length; i += 1 {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func random(num int) int {
	println("Ejercicio 7: numero alreatorio")

	numeroRandom := rand.Intn(num)

	return numeroRandom
}
