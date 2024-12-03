package files

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

// GenerarTabla genera la tabla de multiplicar de un número y la guarda en un archivo .txt.
func GenerarTabla(numero int) {
	// Nombre del archivo a generar
	nombreArchivo := fmt.Sprintf("tabla_%d.txt", numero)

	// Crear archivo
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Printf("Error al crear el archivo: %v\n", err)
		return
	}
	defer archivo.Close()

	// Escribir en el archivo
	writer := bufio.NewWriter(archivo)
	_, _ = writer.WriteString(fmt.Sprintf("Tabla de multiplicar del %d:\n", numero))
	for i := 1; i <= 10; i++ {
		_, _ = writer.WriteString(fmt.Sprintf("%d x %d = %d\n", numero, i, numero*i))
	}

	// Guardar cambios en el archivo
	if err := writer.Flush(); err != nil {
		fmt.Printf("Error al guardar el archivo: %v\n", err)
		return
	}

	fmt.Printf("Tabla generada y guardada en: %s\n", nombreArchivo)
}
