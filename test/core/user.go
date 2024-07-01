package core

import (
	"fmt"
	"regexp"
)

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

type User struct {
	Name  string
	Age   int
	Email string
}

func ValidateUser(u User) (isValid bool, err error) {
	if u.Name == "" {
		return false, fmt.Errorf("el nombre no puede estar vacío")
	}
	if !isValidEmail(u.Email) {
		return false, fmt.Errorf("el correo no es válido")
	}
	if u.Age < 18 {
		return false, fmt.Errorf("debe ser mayor de 18")

	}
	return true, nil
}
