//11
//Api que devuelva la longitud de un string pasado por JSON en body

package main

import (
	"github.com/gin-gonic/gin"
)

func postLongitud(ctx *gin.Context) {
	var userString string
	err := ctx.BindJSON(&userString)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error": "Input must be a string"})
		return
	}
	longitud := len(userString)
	ctx.IndentedJSON(200, gin.H{"Longitud del string": longitud})

}

func main() {
	router := gin.Default()

	router.POST("/longitud", postLongitud)

	router.Run("localhost:8080")
}
