package ejercicios

import (
	"strconv"
)

// ConvertirYEvaluar es una función que convierte un string a un entero
// y devuelve el entero y un mensaje según su valor.
func ConvertirYEvaluar(input string) (int, string) {
	valor, err := strconv.Atoi(input)
	if err != nil {
		return 0, "Error: No es un número válido" + err.Error()
	}

	if valor > 100 {
		return valor, "Es mayor a 100"
	}
	return valor, "Es menor a 100"
}
