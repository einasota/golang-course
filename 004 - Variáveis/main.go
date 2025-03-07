// Metodos de Escrita de Variaveis
package main

import "fmt"

func main() {
  var variavel string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel, variavel2)
  var {
    variavel3 string = "lalala"
    variavel4 string = "lalala"
  }

  fmt.Println(variavel3, variavel4)

  variavel5, variavel6 := "Variavel5", "Variavel6"
  fmt.Println(variavel5, variavel6)

  cont constante1 string = "Constante 1"
  fmt.Println(constante1)
}
