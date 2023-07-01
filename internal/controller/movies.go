package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/splekhanov/go-simple-crud/internal/database"
	"github.com/splekhanov/go-simple-crud/internal/model"
	"github.com/swaggo/swag/example/celler/httputil"
)

// ShowAccount godoc
// @Summary      Creates a movie
// @Description  create a new movie
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        movie	body model.Movie	true	"Add movie"
// @Success      200	{object}	model.Movie
// @Failure		 400	{object}	httputil.HTTPError{400}
// @Failure		 404	{object}	httputil.HTTPError{404}
// @Failure		 500	{object}	httputil.HTTPError{500}
// @Router       /movies [post]
func CreateMovie(c *gin.Context) {
	var movie *model.Movie
	err := c.ShouldBind(&movie)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	res := database.DB.Create(movie)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating a book",
		})
		return
	}
	c.JSON(http.StatusOK, movie)
	return
}
