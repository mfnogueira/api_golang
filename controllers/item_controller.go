// controllers/item_controller.go
package controllers

import (
	"api_golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var requestBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	// Faz o binding do JSON enviado no corpo da requisição
	if err := c.BindJSON(&requestBody); err != nil {
		utils.LogError("Erro ao decodificar o corpo da requisição")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Valida o email
	if !utils.ValidateEmail(requestBody.Email) {
		utils.LogError("Email inválido fornecido: " + requestBody.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email inválido"})
		return
	}

	// Sucesso ao criar o item
	utils.LogInfo("Item criado com sucesso: " + requestBody.Name)
	c.JSON(http.StatusCreated, gin.H{"message": "Item criado com sucesso", "item": requestBody})
}
