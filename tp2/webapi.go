package main

import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

func init(){
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=")
}

func main(){

}