//13
//Hacer una web api que reciba una palabra, y que con una librería que haga yo, use la función que define en la librería para devolver la longitud de la palabra

package main

import (
	"api-libreria-contar-letras/lib"

	"github.com/gin-gonic/gin"
)

func postContarLetras(ctx *gin.Context){
	var palabra string
	err := ctx.BindJSON(&palabra)

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error":"Invalid input"})
		return
	}

	cantidad := lib.ContarLetras(palabra)

	ctx.IndentedJSON(200, cantidad)
}

func main(){
	router := gin.Default()

	router.POST("/contarLetras", postContarLetras)

	router.Run("localhost:8080")
}