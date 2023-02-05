/*Practica 01: Suma de dos números enteros

Enunciado: dado dos números enteros, hallar la suma.

Análisis: para la solución de este problema, se requiere que el usuario ingrese dos números enteros y el sistema realice el cálculo respectivo para hallar la suma.
*/

package main

// `fmt` es un paquete que contiene funciones para formatear e imprimir cadenas.
import "fmt"

func main() {
	//Definicion de variables: var nombreVariable tipo
    var numero1  int
    var numero2 int
    var suma int


		//Imprimir por pantalla y leer desde teclado
    fmt.Print("Ingrese el primer número:")
    fmt.Scan(&numero1)
    fmt.Print("Ingrese el segundo numero:")
    fmt.Scan(&numero2)

		suma = numero1 + numero2

		fmt.Print("La suma de los numeros introducidos es ", suma)
}
