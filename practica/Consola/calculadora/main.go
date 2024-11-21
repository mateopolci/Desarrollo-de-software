// Calculadora suma resta multiplicacion division

package main

import "fmt"

func main() {
	var (
		num1           float64
		num2           float64
		res            float64
		operacion      string
	)

	for {
		fmt.Printf("\nQue operacion desea realizar?\n")
		fmt.Printf("Ingresar 'suma' 'resta' 'multiplicacion' o 'division': ")
		fmt.Scanln(&operacion)

		switch operacion {
		case "suma":
			fmt.Printf("Ingrese primer numero a sumar: ")
			fmt.Scanln(&num1)
			fmt.Printf("Ingrese segundo numero a sumar: ")
			fmt.Scanln(&num2)
			res = num1 + num2
			fmt.Printf("La suma da: %f", res)
		case "resta":
			fmt.Printf("Ingrese primer numero a restar: ")
			fmt.Scanln(&num1)
			fmt.Printf("Ingrese segundo numero a restar: ")
			fmt.Scanln(&num2)
			res = num1 - num2
			fmt.Printf("La resta da: %f", res)
		case "multiplicacion":
			fmt.Printf("Ingrese primer numero a multiplicar: ")
			fmt.Scanln(&num1)
			fmt.Printf("Ingrese segundo numero a multiplicar: ")
			fmt.Scanln(&num2)
			res = num1 * num2
			fmt.Printf("La multiplicacion da: %f", res)
		case "division":
			fmt.Printf("Ingrese primer numero a dividir: ")
			fmt.Scanln(&num1)
			fmt.Printf("Ingrese segundo numero a dividir: ")
			fmt.Scanln(&num2)
			res = num1 / num2
			fmt.Printf("La division da: %f", res)
		}
	}
}
