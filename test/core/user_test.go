package core

import (
	"fmt"
	"testing"
)

func TestValidate_TestTables(t *testing.T) {
	tests := map[string]struct {
		user            User
		expectedIsValue bool
		expectedError   error
	}{
		"EmptyName": {
			user: User{
				Name:  "",
				Age:   18,
				Email: "a@a.com",
			},
			expectedIsValue: false,
			expectedError:   fmt.Errorf("el nombre no puede estar vacío"),
		},
		"EmptyEmail": {
			user: User{
				Name:  "Valentin Zoia",
				Age:   18,
				Email: "invalidgmail.com",
			},
			expectedIsValue: false,
			expectedError:   fmt.Errorf("el correo no es válido"),
		},

		"EmptyAge": {
			user: User{
				Name:  "Valentin Zoia",
				Age:   12,
				Email: "invalid@gmail.com",
			},
			expectedIsValue: false,
			expectedError:   fmt.Errorf("debe ser mayor de 18"),
		},

		"ValidUser": {
			user: User{
				Name:  "Valentin Zoia",
				Age:   19,
				Email: "anashe@gmail.com",
			},
			expectedIsValue: true,
			expectedError:   nil,
		},
	}

	for testName, subTest := range tests {
		t.Run(testName, func(t *testing.T) {
			user := subTest.user
			//por motivos de error. el validuser lo hacemos por separado
			if testName == "ValidUser" {
				//isValid debe ser true

				isValid, err := ValidateUser(user)
				if isValid != subTest.expectedIsValue {
					t.Error("El usuario debe ser valido")
				}
				if err != subTest.expectedError {
					t.Errorf("No se esperaba un eror. Se obtuvo: %s", err.Error())
				}
				return
			}

			// Asegúrate de inicializar correctamente la variable subTest.user

			isValid, err := ValidateUser(user)
			if isValid != subTest.expectedIsValue {
				t.Errorf("Se esperaba un valor de tipo: '%t' pero se obtuvo: '%t", subTest.expectedIsValue, isValid)
			}
			if err.Error() != subTest.expectedError.Error() {
				t.Errorf("Se esperaba un error de tipo: '%s' pero se obtuvo: '%s", subTest.expectedError.Error(), err.Error())
			}

		})
	}

}
