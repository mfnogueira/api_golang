// routes/routes.go
package routes

import (
	"api_golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rota para criar um item
	router.POST("/items", controllers.CreateItem)

	// Rota para listar todos os itens
	router.GET("/items", controllers.ListItems)
}
