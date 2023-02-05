/*
	Se ingresa por teclado un número positivo de uno o dos dígitos (1..99)
	mostrar un mensaje indicando si el número tiene uno o dos dígitos.
	(Tener en cuenta que condición debe cumplirse para tener
		dos dígitos un número entero)
*/

package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	if numero <=9 {
		fmt.Print("El número tiene un solo dígito. ")
	}else{
		if numero <=99{
			fmt.Print("El número tinen dos digitos")
		}else{
			fmt.Print("El número tinen más de dos digitos")
		}
	}

}
