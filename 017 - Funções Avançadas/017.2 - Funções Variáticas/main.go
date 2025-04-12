//Função Variáticas
package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numeros := range numeros {
		total += numeros
	}
	return total
}

func escrever(texto string, numeros ...int)  {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totaldaSoma := soma(1,2,35,6,743,5234,66357,74524)
	fmt.Println(totaldaSoma)
	escrever("Olá Mundo", 10, 2, 4, 6, 7, 8)
}
