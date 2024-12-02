package utils

import (
	"log"
	"os"
)

// LogInfo escreve uma mensagem de info no arquivo de log
func LogInfo(message string) {
	log.SetOutput(os.Stdout)
	log.Println("INFO:", message)
}

// LogError escreve uma mensagem de erro no arquivo de log
func LogError(message string) {
	log.SetOutput(os.Stderr)
	log.Println("ERROR:", message)
}
