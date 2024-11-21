package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Persona struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

var personas = []Persona{
	{
		Id:       1,
		Nombre:   "Mateo",
		Apellido: "Polci",
	},
	{
		Id:       2,
		Nombre:   "Oriana",
		Apellido: "Intendente",
	},
	{
		Id:       3,
		Nombre:   "Felipe",
		Apellido: "Polci",
	},
}

var db *gorm.DB

func initBD() {
	var err error

	db, err = gorm.Open(sqlite.Open("usuarios.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Persona{})
	if err != nil {
		panic(err)
	}

	for _, persona := range personas {
		db.FirstOrCreate(&persona, persona)
	}
}

func getPersonaId(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	var persona Persona

	db.First(&persona, id)

	ctx.JSON(200, persona)
}

func main() {
	initBD()

	router := gin.Default()

	router.GET("/persona/:id", getPersonaId)

	router.Run("localhost:8080")
}
