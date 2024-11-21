//14
//Hacer un endpoint que reciba un número como parámetro, mostrar la cantidad de datos de una lista como diga el número

package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type persona struct{
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Dni int `json:"dni"`
}

var personas = []persona {
	{
		Nombre: "Mateo",
		Apellido: "Polci",
		Dni: 43903414,
	},
	{
		Nombre: "Mateo",
		Apellido: "Polci",
		Dni: 43903414,
	},
	{
		Nombre: "Mateo",
		Apellido: "Polci",
		Dni: 43903414,
	},
	{
		Nombre: "Mateo",
		Apellido: "Polci",
		Dni: 43903414,
	},
	{
		Nombre: "Mateo",
		Apellido: "Polci",
		Dni: 43903414,
	},
}

func getCantidadElementos(ctx *gin.Context){
	cantidad, _ := strconv.Atoi(ctx.Param("cantidad"))

	for i := 0; i < cantidad && i < len(personas); i++ {
		ctx.IndentedJSON(200, personas[i])
	}

}

func main() {
	router := gin.Default()

	router.GET("/datos/:cantidad", getCantidadElementos)

	router.Run("localhost:8080")
}
