package main

import "fmt"

func main() {
	var primerPalabra string
	var segundaPalabra string
	fmt.Printf("Ingrese la primera palabra: ")
	fmt.Scanln(&primerPalabra)
	fmt.Printf("Ingrese la segunda palabra: ")
	fmt.Scanln(&segundaPalabra)

	fmt.Println(primerPalabra+segundaPalabra)
}