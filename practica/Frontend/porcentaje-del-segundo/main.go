// Desarrollar una aplicacion donde dados dos valores el primero sea el porcentaje del segundo

package main

import "fmt"

func main() {
	var (
		num1 float64
		num2 float64
		res float64
	)

	fmt.Printf("Ingresar primer valor: ")
	fmt.Scanln(&num1)
	fmt.Printf("Ingresar segundo valor: ")
	fmt.Scanln(&num2)

	res = (num1 * num2) / 100

	fmt.Printf("El %f porciento de %f es %f porciento.",num1,num2,res)
}
