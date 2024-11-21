package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

var users = []user{
	{
		Id:        1,
		Firstname: "Mateo",
		Lastname:  "Polci",
		User:      "mat",
		Password:  "1234",
	},
	{
		Id:        2,
		Firstname: "Oriana",
		Lastname:  "Monaldi",
		User:      "ori",
		Password:  "4321",
	},
	{
		Id:        3,
		Firstname: "Cooper",
		Lastname:  "Polci",
		User:      "coopy",
		Password:  "0000",
	},
	{
		Id:        4,
		Firstname: "Cloe",
		Lastname:  "Monaldi",
		User:      "michi",
		Password:  "0000",
	},
}

func deleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error": "Input incorrecto"})
		return
	}

	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			ctx.IndentedJSON(200, gin.H{"Ok": "Usuario eliminado"})
			return
		}
	}

	ctx.IndentedJSON(404, gin.H{"Error": "Usuario no encontrado"})
}

func modifyUser(ctx *gin.Context) {
	id, err1 := strconv.Atoi(ctx.Param("id"))

	if err1 != nil {
		ctx.IndentedJSON(400, gin.H{"Error": "Input de id incorrecto"})
		return
	}

	var usuarioModificado user

	err2 := ctx.BindJSON(&usuarioModificado)

	if err2 != nil {
		ctx.IndentedJSON(400, gin.H{"Error": "Input de datos de ususario incorrecto"})
		return
	}

	for i, u := range users {
		if u.Id == id {
			users[i] = usuarioModificado
			ctx.IndentedJSON(200, gin.H{"Ok": "Usuario modificado"})
			return
		}
	}

	ctx.IndentedJSON(404, gin.H{"Error": "Usuario no encontrado"})
}

func getAll(ctx *gin.Context) {
	ctx.IndentedJSON(200, users)
}

func main() {
	router := gin.Default()

	router.DELETE("/delete/:id", deleteUser)
	router.GET("/getAll", getAll)
	router.PUT("/modify/:id", modifyUser)

	router.Run("localhost:8080")
}
