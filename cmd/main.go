package main

import (
	"log"
	"rest/models"
	"rest/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine

func CreateRouter() {
	router = gin.Default()
	router.GET("/todos", routes.Getlist)
	router.POST("/todo", routes.Createitem)
	router.PUT("/todo/:id", routes.Updatelist)
	router.DELETE("/todo/:id", routes.Deleteitem)
}

func main() {
	err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	CreateRouter()

	log.Fatal(router.Run(":8080"))
}
