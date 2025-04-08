//Funções
package main
import "fmt"

func somar(n1 int8, n2 int8) int8 { //Os paramentros e seus tipos vão dentro do parenteses e o tipo do retorno vai após os parenteses e antes da chave.
	return n1 + n2
}
																			//Dois retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10,20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	_, resultadoSubtracao := calculosMatematicos(10, 15) // Para escolher apenas um dos retornos da função!
	fmt.Println(resultadoSubtracao)
}
