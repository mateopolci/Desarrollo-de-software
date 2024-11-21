//3
//Api que reporte una lista de nombres (get) por ID

package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

type nombre struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

var nombres = []nombre  {
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
		Nombre: "Cooper",
	},
	{
		Id : 4,
		Nombre: "Cloe",
	},
}

func getNombre(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error 400":"El id ingresado no pudo ser convertido a numero entero"})
		return
	}
	
	for _, i := range nombres {
		if i.Id == id {
			ctx.IndentedJSON(200, i.Nombre)
			return
		}
	}
	
	ctx.IndentedJSON(404, gin.H{"Error 404":"El id ingresado no se pudo encontrar"})
}

func main() {

	router := gin.Default()

	router.GET("/:id", getNombre)

	router.Run("localhost:8080")
}