package controller

import (
	"github.com/gin-gonic/gin"
	"golang-api/model"
	"net/http"
)

func AuthorController(router *gin.Engine) {
	path := "/api/authors"

	router.GET(path, func(c *gin.Context) {
		authors := model.Authors
		c.IndentedJSON(http.StatusOK, authors)
	})

	router.GET(path+"/:id", func(c *gin.Context) {
		authorId := c.Param("id")
		var author model.Author
		for _, previewAuthor := range model.Authors {
			if previewAuthor.ID == authorId {
				author = previewAuthor
			}
		}
		if (model.Author{}) == author {
			c.IndentedJSON(http.StatusNotFound, nil)
		} else {
			c.IndentedJSON(http.StatusOK, author)
		}
	})
}
