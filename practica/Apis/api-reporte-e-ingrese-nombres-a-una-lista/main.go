// 7
// Api que reporte e ingrese nombres a una lista (GET y POST)
package main

import "github.com/gin-gonic/gin"

type nombre struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

var nombres = []nombre{
	{
		Id:     1,
		Nombre: "Mateo",
	},
	{
		Id:     2,
		Nombre: "Oriana",
	},
	{
		Id:     3,
		Nombre: "Felipe",
	},
}

func postNombre(ctx *gin.Context) {
	var nuevoNombre nombre
	err := ctx.BindJSON(&nuevoNombre)

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error": "Invalid input"})
		return
	}

	nombres = append(nombres, nuevoNombre)
	ctx.IndentedJSON(200, gin.H{"Ok": "Nombre cargado"})
	
}

func getNombres(ctx *gin.Context) {
	ctx.IndentedJSON(200, nombres)
}

func main() {
	router := gin.Default()

	router.GET("/nombres", getNombres)
	router.POST("/nombre", postNombre)

	router.Run("localhost:8080")
}
