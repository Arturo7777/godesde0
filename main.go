package main

import (
	"github.com/Arturo/godesde0/files"
)

//import "github.com/Arturo/godesde0/ejercicios"

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "Linux." || os == "OS X" {
		fmt.Println("Esto es Windows")
	} else {
		fmt.Println("Esto es Windows")
	}
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwing")
	default:
		fmt.Printf("%s \n", os)
	}

	// Llamada a la función ConvertirYEvaluar
	numero, mensaje := ejercicios.ConvertirYEvaluar("150")
	fmt.Printf("Número: %d, Mensaje: %s\n", numero, mensaje)

	numero, mensaje = ejercicios.ConvertirYEvaluar("50")
	fmt.Printf("Número: %d, Mensaje: %s\n", numero, mensaje)

	teclado.IngresoNumero()*/
	numero := files.LeerNumero()
	files.GenerarTabla(numero)

}
