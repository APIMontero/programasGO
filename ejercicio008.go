/*
	Se debe desarrollar un programa que pida el ingreso del
	precio de un artículo y la cantidad que lleva el cliente.
	Mostrar lo que debe abonar el comprador.
	Definir una variable float32 para el precio del artículo.
*/

package main
import "fmt"

func main() {
	var cantidad int
	var precio float32

	fmt.Println("Ingrese el precio del artículo: ")
	fmt.Scan(&precio)

	fmt.Print("Ingrese la cantidad de producto: ")
	fmt.Scan(&cantidad)

	fmt.Print("El comprador debe abonar: $",(precio *float32( cantidad)))
}
