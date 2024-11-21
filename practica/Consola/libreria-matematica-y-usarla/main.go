package main

import (
	"fmt"
	libMat "libreria-matematica/library"
)

func main() {
	
	var (
		suma1           float64
		suma2           float64
		resta1          float64
		resta2          float64
		multiplicacion1 float64
		multiplicacion2 float64
		division1       float64
		division2       float64
	)

	fmt.Printf("Ingrese el primer numero para sumar: ")
	fmt.Scanln(&suma1)
	fmt.Printf("Ingrese el segundo numero para sumar: ")
	fmt.Scanln(&suma2)
	fmt.Printf("Suma: %f\n\n", libMat.Sumar(suma1, suma2))

	fmt.Printf("Ingrese un numero para restarle: ")
	fmt.Scanln(&resta1)
	fmt.Printf("Ingrese el numero a restarle: ")
	fmt.Scanln(&resta2)
	fmt.Printf("Resta: %f\n\n", libMat.Restar(resta1, resta2))
	
	fmt.Printf("Ingrese un numero a multiplicar: ")
	fmt.Scanln(&multiplicacion1)
	fmt.Printf("Ingrese otro numero para multiplicar: ")
	fmt.Scanln(&multiplicacion2)
	fmt.Printf("Multiplicacion: %f\n\n", libMat.Multiplicar(multiplicacion1, multiplicacion2))
	
	fmt.Printf("Ingrese un numero para dividir: ")
	fmt.Scanln(&division1)
	fmt.Printf("Ingrese el numero a dividirle: ")
	fmt.Scanln(&division2)
	fmt.Printf("Division: %f\n\n", libMat.Dividir(division1, division2))
}
