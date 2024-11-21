package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Album struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var albums = []Album{
	{
		Id:    1,
		Name:  "Fly me to the moon",
		Price: 24.39,
	},
	{
		Id:    2,
		Name:  "A sky full of stars",
		Price: 12.45,
	},
	{
		Id:    3,
		Name:  "Eyes on me",
		Price: 7.34,
	},
	{
		Id:    4,
		Name:  "The way you look tonight",
		Price: 14.56,
	},
}

var db *gorm.DB

func initBD() {
	var err error

	//Abrir la DB
	db, err = gorm.Open(sqlite.Open("albums.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Migrar la estructura
	err = db.AutoMigrate(&Album{})
	if err != nil {
		panic(err)
	}

	//Cargar los albumes
	for _, album := range albums {
		db.FirstOrCreate(&album, album)
	}
}

//GET de todos los albumes
func getAlbums(ctx *gin.Context) {
	var albums []Album
	db.Find(&albums)
	ctx.IndentedJSON(200, albums)
}

//GET de un album por id
func getAlbumById(ctx *gin.Context) {
	var album Album
	
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	db.First(&album, id)
	
	ctx.IndentedJSON(200, album)
}

//POST de un album
func postAlbum(ctx *gin.Context) {
	var album Album
	err := ctx.BindJSON(&album) 

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error":"Invalid request"})
		return
	}
	
	db.FirstOrCreate(&album, album)
	ctx.IndentedJSON(200, gin.H{"Ok":"Album añadido"})
}

func modifyAlbumsById(ctx *gin.Context) {
	var album Album
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error":"Invalid ID"})
		return
	}

	err = ctx.BindJSON(&album)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error":"Invalid request"})
		return
	}
	
	db.Model(&Album{}).Where("id = ?", id).Updates(album)
	ctx.IndentedJSON(200, gin.H{"Ok":"Album updated"})
}

func deleteAlbumById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(400, gin.H{"Error":"Invalid ID"})
		return
	}
	
	db.Delete(&Album{}, id)
	ctx.IndentedJSON(400, gin.H{"Ok":"Album eliminado"})
}

func main() {
	initBD()

	router := gin.Default()

	router.GET("albums", getAlbums)
	router.GET("album/:id", getAlbumById)
	router.POST("album", postAlbum)
	router.PUT("album/:id", modifyAlbumsById)
	router.DELETE("album/:id", deleteAlbumById)

	router.Run("localhost:8080")
}
