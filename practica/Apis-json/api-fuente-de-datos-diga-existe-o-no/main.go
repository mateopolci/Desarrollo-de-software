package main

import (
	"encoding/json"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Persona struct {
	Id       int `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

type Personas struct {
	Personas []Persona `json:"personas"`
}

func loadJSON() (Personas, error) {
	var personas Personas

	bytes, err1 := os.ReadFile("personas.json")

	if err1 != nil {
		return personas, nil
	}

	err2 := json.Unmarshal(bytes, &personas)

	if err2 != nil {
		return personas, nil
	}

	return personas, nil
}

func getPersonas(ctx *gin.Context) {
	personas, err := loadJSON()
	
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Server error": "Unable to process request"})
		return
	}
	
	ctx.IndentedJSON(200, personas)
}

func getPersonaId(ctx *gin.Context) {
	id, err1 := strconv.Atoi(ctx.Param("id"))
	
	if err1 != nil {
		ctx.IndentedJSON(400, gin.H{"Server error": "Unable to process request"})
		return
	}
	
	personas, err2 := loadJSON()
	
	if err2 != nil {
		ctx.IndentedJSON(400, gin.H{"Server error": "Unable to process request"})
		return
	}
	
	for _, p := range personas.Personas{
		if p.Id == id {
			ctx.IndentedJSON(200, p)
			return
		}
	}
	
	ctx.IndentedJSON(404, gin.H{"Error": "Unable to find persona with requested ID"})

}

func main() {
	router := gin.Default()

	router.GET("/personas", getPersonas)
	router.GET("/persona/:id", getPersonaId)

	router.Run("localhost:8080")
}
