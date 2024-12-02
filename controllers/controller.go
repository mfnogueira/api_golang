// controllers/controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Função para listar todos os itens
func ListItems(c *gin.Context) {
	// Lógica para listar os itens (pode ser feito com banco de dados, etc)
	items := []string{"Item1", "Item2", "Item3"}
	c.JSON(http.StatusOK, gin.H{"items": items})
}
