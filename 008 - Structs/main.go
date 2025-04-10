// Structs
package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero int8
}

func main(){
	var u usuario
	u.nome = "Davi"
	u.idade = 18
	fmt.Println(u)

	enderecoEx := endereco{"Rua dos Bobos", 0}

	u2 := usuario{"Davi", 21, enderecoEx}
	fmt.Println(u2)

	u3 := usuario{nome: "Davi"}
	fmt.Println(u3)
}