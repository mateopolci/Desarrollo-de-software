//9
//Hacer la suma en una api a través de una biblioteca (2023)

package main

import (
	"github.com/gin-gonic/gin"
	"api-hacer-suma-con-biblioteca/libraries"
)

type input struct {
	Numero1 float64 `json:"numero1"`
	Numero2 float64 `json:"numero2"`
}

func postSuma(ctx *gin.Context) {
	var userInput input

	err := ctx.BindJSON(&userInput)

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error": "Invalid input"})
		return
	}

	res := mate.Sumar(userInput.Numero1, userInput.Numero2)

	ctx.IndentedJSON(200, gin.H{"Resultado": res})

}

func main() {
	router := gin.Default()

	router.POST("/suma", postSuma)

	router.Run("localhost:8080")
}
