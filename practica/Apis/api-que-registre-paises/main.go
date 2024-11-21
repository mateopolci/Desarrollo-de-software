//2
//Realizar una api que registre paises

package main

import (
	"github.com/gin-gonic/gin"
)

type pais struct {
	Id     int `json:"id"`
	Nombre string `json:"nombre"`
}

var paises []pais

func getPaises(ctx *gin.Context) {
	ctx.IndentedJSON(200, paises)
}

func postPais(ctx *gin.Context) {
	var nuevoPais pais

	err := ctx.BindJSON(&nuevoPais)

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error": "Bad request, JSON could not be parsed"})
	}

	paises = append(paises, nuevoPais)
	ctx.IndentedJSON(200, gin.H{"OK": "Pais añadido a la lista"})

}

func main() {
	router := gin.Default()
	router.GET("/paises", getPaises)
	router.POST("/paises", postPais)
	router.Run("localhost:8080")
}
