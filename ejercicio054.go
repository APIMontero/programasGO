/*
uri:"https://www.tutorialesprogramacionya.com/goya/detalleconcepto.php?punto=11&codigo=11&inicio=0",
numero:054,
enunciado:"
	Se cuenta con la siguiente información:
	Las edades de 5 estudiantes del turno mañana.
	Las edades de 6 estudiantes del turno tarde.
	Las edades de 11 estudiantes del turno noche.
	Las edades de cada estudiante deben ingresarse por teclado.
	a) Obtener el promedio de las edades de cada turno (tres promedios)
	b) Imprimir dichos promedios (promedio de cada turno)
	c) Mostrar por pantalla un mensaje que indique cual de los tres turnos tiene un promedio de edades mayor. "
*/

package main
import "fmt"

func promedio(estudiantes float32, turno string){
		var suma float32
		var promedio float32
		var edad   float32
	  fmt.Print("Ingrese los datos para el turno de",turno)
		for f := 1; f <= estudiantes; f++ {
        fmt.Print("Ingrese edad del estudiante # ",f,":")
        fmt.Scan(&edad)
				suma = suma + edad
    }

		promedio=float32(suma) / float32(estudiantes)

    fmt.Println("El promedio del turno de ",turno," es:", promedio)

		return promedio
}

func main() {
	/*
	Las edades de los estudiantes del cada  turno y sus promedios
*/

var promedio5 , promedio6,  promedio11, mayor float32
promedio5=promedio(5,"Mañana")
promedio6= promedio(6,"Tarde")
promedio11 = promedio(11,"Noche")
mayor  = promedio5
turno := " "

if promedio5 > mayor{
	mayor=promedio5
	turno="Turno de Mañana"
} else if promedio6 > mayor{
	mayor=promedio6
  turno="Turno de Tarde"
}else if promedio11 > mayor{
	mayor=promedio11
	turno="Turno de Noche"
}

fmt.Print("El promedio mayor fue del ", turno, " con promedio de ", mayor, " estudiantes.")
}
