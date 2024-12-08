package ejer_interfaces

import (
	"fmt"

	"github.com/Arturo/godesde0/interfacess"
)

func HumanosRespirando(hu interfacess.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
