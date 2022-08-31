package main

import (
	"github.com/gin-gonic/gin"
	"golang-api/controller"
)

func main() {
	router := gin.Default()
	controller.AuthorController(router)
	router.Run("localhost:8082")
}
