//4
//Api que reporte una lista de nombres (get)

package main

import "github.com/gin-gonic/gin"

func getNombres(ctx *gin.Context) {
	ctx.IndentedJSON(200, nombres)
}

type nombre struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

var nombres = []nombre{
	{
		Id : 1,
		Nombre: "Mateo",
	},
	{
		Id : 2,
		Nombre: "Oriana",
	},
	{
		Id : 3,
		Nombre: "Cloe",
	},
}

func main() {

	router := gin.Default()

	router.GET("/nombres", getNombres)

	router.Run("localhost:8080")
}
