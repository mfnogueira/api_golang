package utils

import (
	"log"
	"regexp"
)

// ValidateEmail valida o formato de um email
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if re.MatchString(email) {
		return true
	}
	log.Println("Erro: Email inválido")
	return false
}
