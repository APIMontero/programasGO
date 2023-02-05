/*
	Realizar un programa que solicite por teclado cuatro valores numéricos enteros
	e informar su suma y promedio.
*/

package main
import "fmt"

func main() {
	var uno,dos,tres,cuatro, suma  int
	var promedio float32

	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&uno)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&dos)
	fmt.Print("Ingrese el tercer número: ")
	fmt.Scan(&tres)
	fmt.Print("Ingrese el cuarto número: ")
	fmt.Scan(&cuatro)

	suma = uno+dos+tres+cuatro
	promedio=float32(suma) / 4.0

	fmt.Print("La suma de todos los numeros es: ", suma)
	fmt.Print("\nEl promedio de todos los números es: ", promedio)
}
