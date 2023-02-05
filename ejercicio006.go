
/* Escribir un programa en el cual se ingresen cuatro números,
   calcular e informar la suma de los dos primeros y el producto
   del tercero y el cuarto. */

package main
import "fmt"

func main() {
	var uno,dos,tres,cuatro, suma, producto int

	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&uno)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&dos)
	fmt.Print("Ingrese el tercer número: ")
	fmt.Scan(&tres)
	fmt.Print("Ingrese el cuarto número: ")
	fmt.Scan(&cuatro)

	suma = uno+dos
	producto=tres*cuatro

	fmt.Print("La suma de ", uno, " con ", dos, " es: ", suma)
	fmt.Print("\nEl producto de ", tres, " con ", cuatro, " es: ", producto)
}
