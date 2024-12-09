package main

import (
	"fmt"

	"github.com/Arturo/godesde0/goroutines"
)

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
	//numero := files.LeerNumerob()
	//files.GenerarTablab(numero)

	//funciones.Calculos()
	//funciones.Exponencia(2)
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)/-*/

	//d.EjemploPanic()
	canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Arturo Rodriguez", canal1)
	fmt.Println("Estoy aqui")

	<-canal1

}
