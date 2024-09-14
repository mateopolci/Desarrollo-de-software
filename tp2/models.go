package main

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Name string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	gorm.Model
	Title string `json:"title"`
	AuthorID uint `json:"author_id"`
	Author Author `json:"author"`
}