package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   // slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  // slice creado desde un vector, desde la posicion 0 hasta la 5
	porcion3 := arreglo[1:4] // slice creado desde un vector, desde la posicion 1 hasta la 4

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
