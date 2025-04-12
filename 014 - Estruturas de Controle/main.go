//Estruturas de Controle
package main

import "fmt"

func main()  {
	numero := 10
	if numero > 15 {
		fmt.Println("Maior de 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}
	//If init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que zero")
	} else {
		fmt.Println("Está entre 0 e -10")
	}
}
