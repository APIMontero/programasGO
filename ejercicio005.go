/* Realizar la carga del lado de un cuadrado,
mostrar por pantalla el perímetro del mismo
(El perímetro de un cuadrado se calcula
	multiplicando el valor del lado por cuatro) */

package main
import "fmt"

func main() {

	var ladoCuadrado int
	var perimetroCuadrado int

  fmt.Print("Ingrese el lado del cuadrado:")
  fmt.Scan(&ladoCuadrado)
  perimetroCuadrado = ladoCuadrado * 4
  fmt.Print("El perímetro de un cuadrado de lado ",ladoCuadrado," es: ", perimetroCuadrado)

	fmt.Print("\n")
}
