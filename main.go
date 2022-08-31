package main

import (
	"authors-api/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controller.AuthorController(router)
	router.Run("localhost:8082")
}
