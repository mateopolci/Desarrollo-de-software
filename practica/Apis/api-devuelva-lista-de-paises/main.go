//5
//Api que devuelve una lista de paises (get)

package main

import "github.com/gin-gonic/gin"

type pais struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Poblacion int `json:"poblacion"`
}

var paises = []pais{
	{
		Id : 1,
		Nombre : "Argentina",
		Poblacion : 50000000,
	},
	{
		Id : 2,
		Nombre : "Brasil",
		Poblacion : 90000000,
	},
	{
		Id : 3,
		Nombre : "Uruguay",
		Poblacion : 10000000,
	},
}

func listaPaises(ctx *gin.Context){
	ctx.IndentedJSON(200,paises)
}

func main(){
	router := gin.Default()

	router.GET("/paises",listaPaises)

	router.Run("localhost:8080")
}