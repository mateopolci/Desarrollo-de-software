package main

import (
	"api-hacer-multiplicacion-con-biblioteca/library"
	"github.com/gin-gonic/gin"
)

type input struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

func postMultiplicar(ctx *gin.Context){
	var userInput input
	err := ctx.BindJSON(&userInput)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error":"Invalid body structure, could not parse JSON to object"})
		return
	}
	res := library.Multiplicar(userInput.Num1, userInput.Num2)

	ctx.IndentedJSON(200, gin.H{"Resultado": res})
}

func main(){
	router:=gin.Default()

	router.POST("/multiplicar", postMultiplicar)

	router.Run("localhost:8080")
}