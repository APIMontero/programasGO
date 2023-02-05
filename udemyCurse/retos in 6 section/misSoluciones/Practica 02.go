/*Practica 02: Calcular cociente y el Residuo de dos números enteros

Enunciado: halar el cociente y el residuo (resto) de dos números enteros.

Análisis: para la solución de este problema, se requiere que el usuario ingrese dos números enteros y el sistema realice el cálculo respectivo para hallar el cociente y residuo.
*/

package main

// `fmt` es un paquete que contiene funciones para formatear e imprimir cadenas.
import "fmt"

func main() {
	//Definicion de variables: var nombreVariable tipo
    var numero1  int
    var numero2 int
    var cuociente int
		var residuo int
    		//Imprimir por pantalla y leer desde teclado
    fmt.Print("Ingrese el primer número:")
    fmt.Scan(&numero1)
    fmt.Print("Ingrese el segundo numero:")
    fmt.Scan(&numero2)

		cuociente = numero1 / numero2
		residuo = numero1 % numero2

		fmt.Print("cuociente de los numeros introducidos es ", cuociente, " y el residuo es ", residuo)
}
