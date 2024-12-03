package files

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// LeerNumero pide un número al usuario, validando si hay errores en la entrada.
func LeerNumerob() int {
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

// GenerarTabla genera la tabla de multiplicar de un número y la guarda en un archivo .txt en modo append.
func GenerarTablab(numero int) {
	// Definir la ruta del archivo
	rutaDirectorio := "./files/txt/"
	rutaArchivo := filepath.Join(rutaDirectorio, "tabla.txt")

	// Crear los directorios si no existen
	if err := os.MkdirAll(rutaDirectorio, os.ModePerm); err != nil {
		fmt.Printf("Error al crear directorios: %v\n", err)
		return
	}

	// Abrir el archivo en modo append o crearlo si no existe
	archivo, err := os.OpenFile(rutaArchivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return
	}
	defer archivo.Close()

	// Escribir la tabla de multiplicar en el archivo
	writer := bufio.NewWriter(archivo)
	_, _ = writer.WriteString(fmt.Sprintf("Tabla de multiplicar del %d:\n", numero))
	for i := 1; i <= 10; i++ {
		_, _ = writer.WriteString(fmt.Sprintf("%d x %d = %d\n", numero, i, numero*i))
	}
	_, _ = writer.WriteString("\n") // Añadir una línea en blanco al final

	// Guardar cambios en el archivo
	if err := writer.Flush(); err != nil {
		fmt.Printf("Error al guardar el archivo: %v\n", err)
		return
	}

	fmt.Printf("Tabla de multiplicar del %d guardada correctamente en: %s\n", numero, rutaArchivo)
}
