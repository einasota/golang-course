// Funções com ponteiros
package main

import "fmt"

func inverterSinal(numero int) int { // Parametro por copia
	return numero * -1	
}

func inverterSinalComPonteiro(numero *int) { //Parametro por referencia
	*numero = *numero * -1	
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
