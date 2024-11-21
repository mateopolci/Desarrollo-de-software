// 1
// Crear una base de datos, cargar los datos con usuarios, hacer un endpoint que devuelva usuario por id
package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Usuario struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var db *gorm.DB

func initDB() {
	var err error
	// Abrimos la BDD
	db, err = gorm.Open(sqlite.Open("usuarios.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Creamos la tabla Usuario en la BDD
	err = db.AutoMigrate(&Usuario{})
	if err != nil {
		panic(err)
	}

	// Creamos un array de objetos de tipo Usuario
	usuarios := []Usuario{
		{
			Id: 1, 
			Nombre: "Mateo Polci", 
			Email: "mateopolci@hotmail.com",
		},
		{
			Id: 2, 
			Nombre: "Oriana Intendente Monaldi", 
			Email: "orimonaldi@gmail.com",
		},
		{
			Id: 3, 
			Nombre: "Felipe Polci", 
			Email: "polcifelipe@gmail.com",
		},
	}

	// Inserta el array de usuarios en la base de datos
	for _, usuario := range usuarios {
		db.FirstOrCreate(&usuario, Usuario{Email: usuario.Email})
	}

}

// Endpoint
func getUsuarioId(ctx *gin.Context) {
	var usuario Usuario

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error": "ID inválido"})
		return
	}

	if err := db.First(&usuario, id).Error; err != nil {
		ctx.IndentedJSON(404, gin.H{"Error": "Persona no encontrada"})
		return
	}

	ctx.IndentedJSON(200, usuario)

}


func main() {
	initDB()

	router := gin.Default()

	router.GET("/usuario/:id", getUsuarioId)

	router.Run("localhost:8080")
}
