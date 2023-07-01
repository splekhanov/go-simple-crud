package main

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/splekhanov/go-simple-crud/internal/controller"
	"github.com/splekhanov/go-simple-crud/internal/database"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	v1 := r.Group("/api/v1")

	{
		movies := v1.Group("/movies")
		{
			movies.POST("", controller.CreateMovie)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
