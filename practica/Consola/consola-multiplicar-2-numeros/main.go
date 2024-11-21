// Pedir dos numeros, multiplicarlos y mostralos en consola

package main

import "fmt"

func main(){
	var (
		primerNumero int;
		segundoNumero int;
		producto int;
	)

	fmt.Printf("Ingrese el primer numero: ")
	fmt.Scanln(&primerNumero)
	fmt.Printf("Ingrese el segundo numero: ")
	fmt.Scanln(&segundoNumero)

	producto = primerNumero * segundoNumero

	fmt.Printf("El producto es: %d",producto)
}