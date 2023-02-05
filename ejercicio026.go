/*
	De un operario se conoce su sueldo y los años de antigüedad. Se pide confeccionar un programa que lea los datos de entrada e informe:
	a) Si el sueldo es inferior a 500 y su antigüedad es igual o superior a 10 años, otorgarle un aumento del 20 %, mostrar el sueldo a pagar.
	b)Si el sueldo es inferior a 500 pero su antigüedad es menor a 10 años, otorgarle un aumento de 5 %.
	c) Si el sueldo es mayor o igual a 500 mostrar el sueldo en pantalla sin cambios.
*/

/*
	calculos:
		aumento=condicionado por sueldo y años(condiciones desde a) hasta c) ) float
		sueldo a pagar = sueldo + (sueldo * aumento) float
	Mensajes cadenas
*/
package main
import "fmt"

func main(){
var sueldo, sueldoPagar, aumento float32
var agnos int


fmt.Print("Ingrese los años de antigüedad: ")
fmt.Scan(&agnos)
fmt.Print("Ingrese sueldo: ")
fmt.Scan(&sueldo)
/*
a) Si el sueldo es inferior a 500 y su antigüedad es igual o superior a 10 años,
otorgarle un aumento del 20 %, mostrar el sueldo a pagar.*/
if sueldo <500 && agnos>=10{
	aumento=float32(20)/float32(100)
}
/*
b)Si el sueldo es inferior a 500 pero su antigüedad es menor a 10 años,
 otorgarle un aumento de 5 %.
*/
if sueldo < 500 && agnos < 10{
	aumento=float32(5)/float32(100)
}
/*
c) Si el sueldo es mayor o igual a 500 mostrar el sueldo en pantalla sin cambios.
*/
if sueldo >= 500 {
	aumento=float32(0)
}

sueldoPagar= sueldo + (sueldo * aumento)


fmt.Print("Sueldo a pagar: $", sueldoPagar)
}
