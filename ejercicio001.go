package main

// `fmt` es un paquete que contiene funciones para formatear e imprimir cadenas.
import "fmt"

func main() {
	//Definicion de variables: var nombreVariable tipo
    var horasTrabajadas int
    var costoHora float32
    var sueldo float32

		//Imprimir por pantalla y leer desde teclado
    fmt.Print("Ingrese las horas trabajadas por el empleado:")
    fmt.Scan(&horasTrabajadas)
    fmt.Print("Ingrese el pago por hora:")
    fmt.Scan(&costoHora)
    sueldo=float32(horasTrabajadas) * costoHora
    fmt.Print("El sueldo total del operario es ",sueldo)
}



