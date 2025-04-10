// Herança só que não
package main

import "fmt"

type pessoa struct {
	nome string
	sobremone string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main(){
	p1 := pessoa{"João","Pedro",10,170}

	e1 := estudante{p1,"Engenharia","USP"}

	fmt.Println(e1)
}