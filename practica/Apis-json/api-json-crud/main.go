package main

import (
	"encoding/json"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Persona struct {
	Id       int    `json:"id"`
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
		return personas, err1
	}

	err2 := json.Unmarshal(bytes, &personas)

	if err2 != nil {
		return personas, err2
	}

	return personas, nil
}

func getAll(ctx *gin.Context) {
	personas, err := loadJSON()

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Internal server error": "Error loading JSON"})
		return
	}

	ctx.IndentedJSON(200, personas)
}

func getPersona(ctx *gin.Context) {
	personas, err1 := loadJSON()
	if err1 != nil {
		ctx.IndentedJSON(400, gin.H{"Internal server error": "Error loading JSON"})
		return
	}
	id, err2 := strconv.Atoi(ctx.Param("id"))
	if err2 != nil {
		ctx.IndentedJSON(400, gin.H{"Internal server error": "Error fetching ID from parameter"})
		return
	}
	for _, p := range personas.Personas {
		if p.Id == id {
			ctx.IndentedJSON(200, p)
			return
		}
	}

	ctx.IndentedJSON(404, gin.H{"Error": "ID not found"})
}

func postPersona(ctx *gin.Context) {
	var newPersona Persona

	err1 := ctx.BindJSON(&newPersona)
	if err1 != nil {
		ctx.IndentedJSON(400, gin.H{"Internal server error": "Invalid input"})
		return
	}

	personas, err2 := loadJSON()
	if err2 != nil {
		ctx.IndentedJSON(400, gin.H{"Internal server error": "Error loading JSON"})
		return
	}

	personas.Personas = append(personas.Personas, newPersona)

	bytes, err3 := json.Marshal(personas)
	if err3 != nil {
		ctx.IndentedJSON(400, gin.H{"Internal server error": "Error marshalling JSON"})
		return
	}

	err4 := os.WriteFile("personas.json", bytes, 0644)
	if err4 != nil {
		ctx.IndentedJSON(400, gin.H{"Internal server error": "Error writing JSON"})
		return
	}

	ctx.IndentedJSON(200, gin.H{"Success": "Persona added"})
}

func putPersona(ctx *gin.Context) {
    id, err1 := strconv.Atoi(ctx.Param("id"))
    if err1 != nil {
        ctx.IndentedJSON(400, gin.H{"Internal server error": "Error fetching ID from parameter"})
        return
    }

    var updatedPersona Persona
    err2 := ctx.BindJSON(&updatedPersona)
    if err2 != nil {
        ctx.IndentedJSON(400, gin.H{"Internal server error": "Invalid input"})
        return
    }

    personas, err3 := loadJSON()
    if err3 != nil {
        ctx.IndentedJSON(400, gin.H{"Internal server error": "Error loading JSON"})
        return
    }

    found := false
    for i, p := range personas.Personas {
        if p.Id == id {
            personas.Personas[i] = updatedPersona
            found = true
            break
        }
    }

    if !found {
        ctx.IndentedJSON(404, gin.H{"Error": "ID not found"})
        return
    }

    bytes, err4 := json.Marshal(personas)
    if err4 != nil {
        ctx.IndentedJSON(400, gin.H{"Internal server error": "Error marshalling JSON"})
        return
    }

    err5 := os.WriteFile("personas.json", bytes, 0644)
    if err5 != nil {
        ctx.IndentedJSON(400, gin.H{"Internal server error": "Error writing JSON"})
        return
    }

    ctx.IndentedJSON(200, gin.H{"Ok": "Persona updated"})
}

func deletePersona(ctx *gin.Context) {
    id, err1 := strconv.Atoi(ctx.Param("id"))
    if err1 != nil {
        ctx.IndentedJSON(400, gin.H{"Internal server error": "Error fetching ID from parameter"})
        return
    }

    personas, err2 := loadJSON()
    if err2 != nil {
        ctx.IndentedJSON(400, gin.H{"Internal server error": "Error loading JSON"})
        return
    }

    found := false
    var filteredPersonas []Persona
    for _, p := range personas.Personas {
        if p.Id != id {
            filteredPersonas = append(filteredPersonas, p)
        } else {
            found = true
        }
    }

    if !found {
        ctx.IndentedJSON(404, gin.H{"Error": "ID not found"})
        return
    }

    personas.Personas = filteredPersonas

    bytes, err3 := json.Marshal(personas)
    if err3 != nil {
        ctx.IndentedJSON(400, gin.H{"Internal server error": "Error marshalling JSON"})
        return
    }

    err4 := os.WriteFile("personas.json", bytes, 0644)
    if err4 != nil {
        ctx.IndentedJSON(400, gin.H{"Internal server error": "Error writing JSON"})
        return
    }

    ctx.IndentedJSON(200, gin.H{"Success": "Persona deleted"})
}

func main() {

	router := gin.Default()

	router.GET("/getAll", getAll)
	router.GET("/getPersona/:id", getPersona)
	router.POST("/postPersona", postPersona)
	router.PUT("/putPersona/:id", putPersona)
	router.DELETE("/deletePersona/:id", deletePersona)

	router.Run("localhost:8080")
}