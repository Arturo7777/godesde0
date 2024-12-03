package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// LeerNumero pide un número al usuario, validando si hay errores en la entrada.
func LeerNumero() int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Ingresa un número: ")
		if scanner.Scan() {
			entrada := scanner.Text()
			numero, err := strconv.Atoi(entrada)
			if err != nil {
				fmt.Println("Entrada inválida. Por favor, ingresa un número entero.")
				continue
			}
			return numero
		} else {
			fmt.Println("Error al leer la entrada. Intenta nuevamente.")
		}
	}
}

// GenerarTabla genera y muestra la tabla de multiplicar de un número.
func GenerarTabla(numero int) {
	fmt.Printf("Tabla de multiplicar del %d:\n", numero)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}
}
