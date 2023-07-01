package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/splekhanov/go-simple-crud/internal/database"
)

func CreateMovie(c *gin.Context) {
	var movie *database.Movie
	err := c.ShouldBind(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(movie)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": movie,
	})
	return
}
