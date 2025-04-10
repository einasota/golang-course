// Array Interno
package main

import "fmt"

func main()  {
	slice := make([]float32, 10, 11)
	fmt.Println(slice)
	slice = append(slice, 5)
	slice = append(slice, 6)
	fmt.Println(slice)
	fmt.Println(len(slice)) //Length
	fmt.Println(cap(slice)) //Capacidade
	//Após estourar a capacidade, é criado outro array interno com o dobro da capacidade.
	slice2 := make([]float32, 5)
	fmt.Println(slice2)
	slice = append(slice2,10)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
