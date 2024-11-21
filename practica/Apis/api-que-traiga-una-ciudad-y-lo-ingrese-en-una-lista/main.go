//6
//Api que traiga una ciudad y lo ingrese en una lista (GET y POST)

package main

import "github.com/gin-gonic/gin"

type ciudad struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
}

var ciudades []ciudad

func getCiudad(ctx * gin.Context){
	ctx.IndentedJSON(200, ciudades)
}

func postCiudad(ctx * gin.Context){
	var nuevaCiudad ciudad
	
	err := ctx.BindJSON(&nuevaCiudad)

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error":"El formato de ciudad ingresado no es correcto"})
		return
	}
	
	ciudades = append(ciudades, nuevaCiudad)
	ctx.IndentedJSON(200, gin.H{"Ok":"La ciudad ha sido cargada"})

}

func main(){
	router := gin.Default()

	router.GET("/ciudades", getCiudad)
	router.POST("/ciudad", postCiudad)

	router.Run("localhost:8080")

}