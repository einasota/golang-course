// Operadores
package main

import "fmt"

func main()  {
	//Aritimeticos
	soma := 1 + 2
	subtracao := 1 - 2 
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// Relacionais
	fmt.Println("Maior que",1 > 2)
	fmt.Println("Maior ou igual que", 1 >= 2)
	fmt.Println("Igual", 1 == 2)
	fmt.Println("Menor que", 1 < 2)
	fmt.Println("Menor ou igual que", 1 <= 2)
	fmt.Println("Diferente", 1 != 2)

	// Logicos
	fmt.Println("\n -------------------")
	verdadeiro, falso := true, false
	fmt.Println("AND", verdadeiro && falso)
	fmt.Println("OR", verdadeiro || falso)
	fmt.Println("NOT", !verdadeiro)

	//Unários
	numero3 := 10
	numero3 ++
	numero3 += 15
	numero3 --
	numero3 -= 1
	numero3 *= 3
	numero3 /= 7
	numero3 %= 9
	fmt.Println(numero3)

	//Ternário
	//Não existe operadora Ternário no go. Precisa ser utilizado a estrutura de repetição IF
}