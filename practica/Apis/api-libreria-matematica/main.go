//12
//Crear librería matemática y usar en web api (2023)
package main

import (
	"github.com/gin-gonic/gin"
	"api-libreria-matematica/lib"
)

type input struct{
	Numero1 float64 `json:"numero1"`
	Numero2 float64 `json:"numero2"`
}

func postSumar (ctx *gin.Context){
	var userInput input

	err := ctx.BindJSON(&userInput)

	if err != nil{
		ctx.IndentedJSON(400,gin.H{"Error":"Invalid JSON input"})
		return
	}

	res := lib.Sumar(userInput.Numero1, userInput.Numero2)
	ctx.IndentedJSON(200, gin.H{"Resultado": res})
}

func postRestar (ctx *gin.Context){
	var userInput input

	err := ctx.BindJSON(&userInput)

	if err != nil{
		ctx.IndentedJSON(400,gin.H{"Error":"Invalid JSON input"})
		return
	}

	res := lib.Restar(userInput.Numero1, userInput.Numero2)
	ctx.IndentedJSON(200, gin.H{"Resultado": res})
}

func postMultiplicar (ctx *gin.Context){
	var userInput input

	err := ctx.BindJSON(&userInput)

	if err != nil{
		ctx.IndentedJSON(400,gin.H{"Error":"Invalid JSON input"})
		return
	}

	res := lib.Multiplicar(userInput.Numero1, userInput.Numero2)
	ctx.IndentedJSON(200, gin.H{"Resultado": res})
}

func postDividir (ctx *gin.Context){
	var userInput input

	err := ctx.BindJSON(&userInput)

	if err != nil{
		ctx.IndentedJSON(400,gin.H{"Error":"Invalid JSON input"})
		return
	}

	res := lib.Dividir(userInput.Numero1, userInput.Numero2)
	ctx.IndentedJSON(200, gin.H{"Resultado": res})
}

func main(){
	router := gin.Default()

	router.POST("/sumar", postSumar)
	router.POST("/restar", postRestar)
	router.POST("/multiplicar", postMultiplicar)
	router.POST("/dividir", postDividir)

	router.Run("localhost:8080")
}