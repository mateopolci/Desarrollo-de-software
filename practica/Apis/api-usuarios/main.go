//13
//Hacer un Post que tiene que recibir un JSON que tenga los datos de un usuario(nombre y apellido), guardarlo tipo objeto y luego guardarlo en una lista. Después lo tienes que mostrar en pantalla.

package main

import (
	"github.com/gin-gonic/gin"
)

type usuario struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

var usuarios []usuario

func getUsuarios(ctx *gin.Context) {
	ctx.IndentedJSON(200, usuarios)
}

func postUsuario(ctx *gin.Context) {
	var nuevoUsuario usuario
	err := ctx.BindJSON(&nuevoUsuario)

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error" : "El JSON ingresado es invalido"})
		return
	}
	
	usuarios = append(usuarios, nuevoUsuario)
	ctx.IndentedJSON(200, gin.H{"Ok" : "Usuario añadido a la lista"})
}

func main() {
	router := gin.Default()

	router.GET("/usuarios", getUsuarios)
	router.POST("/usuario", postUsuario)

	router.Run("localhost:8080")
}
