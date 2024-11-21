//8
//Api web para Concatenar dos palabras (2023)

package main

import "github.com/gin-gonic/gin"

type input struct {
	Palabra1 string `json:"palabra1"`
	Palabra2 string `json:"palabra2"`
}

func postConcatenar(ctx *gin.Context) {
	var palabras input
	err := ctx.BindJSON(&palabras)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error": "El input no es un objeto JSON válido"})
		return
	}

	ctx.IndentedJSON(200, gin.H{"Resultado": palabras.Palabra1 + palabras.Palabra2})
}

func main() {
	router := gin.Default()

	router.POST("/concatenar", postConcatenar)

	router.Run("localhost:8080")

}
