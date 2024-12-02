package main

import (
	"api_golang/routes"
	"api_golang/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LogInfo("Iniciando a API em Go...")

	// Inicializa o roteador
	router := gin.Default()

	// Configura as rotas
	routes.SetupRoutes(router)

	// Inicia o servidor
	utils.LogInfo("Servidor iniciado na porta 8080")
	router.Run(":8080")
}
