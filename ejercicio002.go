//Insertaremos codigo erroneo adrede y lo ejecutaremos...
package main
import "fmt"

func main() {
    var lado int
		//CODIGO ORIGINAL: VAR superficie int
    var superficie int//CORREGIDO

    fmt.Print("Ingrese el valor del lado del cuadrado:")
    fmt.Scan(&lado)
    superficie = lado * lado
    fmt.Print("La superficie del cuadrado es:",superficie)
}
/*Resultado:
PS C:\GO\programasGO> go run ejercicio002.go
Ingrese el valor del lado del cuadrado:125
La superficie del cuadrado es:15625
PS C:\GO\programasGO> */
