//1
//CRUD TOTAL

package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

// Definimos globalmente el tipo persona
type persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Dni      int `json:"dni"`
}

// Declaramos e inicializamos globalmente la variable de tipo array de personas
var personas = []persona{
	{
		ID: 1,
		Nombre: "Mateo",
		Apellido: "Polci",
		Dni: 43903414,
	},
	{
		ID: 2,
		Nombre: "Esteban",
		Apellido: "Polci",
		Dni: 18333648,
	},
	{
		ID: 3,
		Nombre: "Eugenia",
		Apellido: "Berlingieri",
		Dni: 17815990,
	},
}

// GET Todas las personas
func getPersonas(ctx *gin.Context){
	ctx.IndentedJSON(200, personas)
}

// GET Persona por ID 
func getPersona(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))
	for _, p := range personas {
		if p.ID == id {
			ctx.IndentedJSON(200, p)
			return
		}
	}
	ctx.IndentedJSON(404, gin.H{"message":"Error 404: Persona no encontrada"})
}

// POST Persona
func postPersona(ctx *gin.Context){
	var nuevaPersona persona
	err := ctx.BindJSON(&nuevaPersona)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message":"Se ha producido un error"})
		return
	}
	personas = append(personas, nuevaPersona)
	ctx.IndentedJSON(200, gin.H{"message":"Persona añadida"})
}

// DELETE Persona por ID
func deletePersona(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))
	for i, p := range personas{
		if p.ID == id {
			personas = append(personas[:i], personas[i+1:]...)
			ctx.IndentedJSON(200, gin.H{"message":"Persona eliminada"})
			return
		}
	}
	ctx.IndentedJSON(404, gin.H{"message":"Error 404: Persona no encontrada"})
}

// PUT Persona por ID
func putPersona(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))
	var personaModificada persona
	err := ctx.BindJSON(&personaModificada)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message":"Formato de JSON incorrecto"})
		return
	}
	for i, p := range personas {
		if p.ID == id {
			personas[i] = personaModificada
			ctx.IndentedJSON(200, gin.H{"message":"Persona modificada"})
			return
		}
	}
}

func main() {
	// Router
	router := gin.Default()

	// Endpoints
	router.GET("/personas", getPersonas)
	router.GET("/persona/:id", getPersona)
	router.POST("/persona", postPersona)
	router.DELETE("/persona/:id", deletePersona)
	router.PUT("/persona/:id", putPersona)

	// URL
	router.Run("localhost:8080")
}
