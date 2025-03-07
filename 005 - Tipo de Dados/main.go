// Tipos de Dados
package main

import "fmt"

func main(){
  //Numeros Inteiros
  //8, 16, 32, 64 bits
  var numero int64 = -100000000000
  fmt.Println(numero)
  //Numeros Inteiros porém sem sinal(-)
  var numero2 uint32 = 10000
  fmt.Println(numero2)
  //rune = int32 (usado para representar caracteres)
  var numero3 rune = 123456
  fmt.Println(numero3)
  //byte = int8
  var numero4 byte = 123
  fmt.Println(numero4)

  //Numeros Reais
  //32 e 64bits
  var numeroReal1 float32 = 123.45
  fmt.Println(numeroReal1)
  var numeroReal2 float64 = 12300000000.45
  fmt.Println(numeroReal2)
  //Na inferência de tipo, o tipo reconhecido será apenas float, porem não é possivel definir a variavel com o tipo float.
  numeroReal3 := 132151232.3432

  //Texto
  var str string = "Texto"
  fmt.Println(str)

  //Inferência de tipo
  str2 := "Texto2"
  
  //O Caracter é declarado utilizando aspas simples(' ') e ele vai ser declarado como int e retornar ovalor do caracter em numero da localização da tabela ASCII

  char := '%'
  fmt.Println(char)

  //Valor 0
  var texto string
  fmt.Println(texto)
  //Todo tipo de dado tem seu valor zero, nesse exemplo ele mostra o uma string vazia, no caso de um numero ele retorna 0, caso se for booleano ele retorna false e caso se for um error ele retorna <nil> que é o nulo no go.

  //Booleano 
  var booleano1 bool = true
  fmt.Println(booleano1)
  //O valor zero é false
  var booleano2 bool
  fmt.Println(booleano2)

  //Erro
  //Server para conseguimos tratar os erros
  var erro error
  fmt.Println(erro)
  //O Valor Zero dele é o <nil>, que significa basicamente nulo.
}
