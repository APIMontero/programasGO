/*
	Un postulante a un empleo, realiza un test de capacitación, se obtuvo la siguiente información: cantidad total de preguntas que se le realizaron y la cantidad de preguntas que contestó correctamente. Se pide confeccionar un programa que ingrese los dos datos por teclado e informe el nivel del mismo según el porcentaje de respuestas correctas que ha obtenido, y sabiendo que:

		Nivel máximo:	Porcentaje>=90%.
		Nivel medio:	Porcentaje>=75% y <90%.
		Nivel regular:	Porcentaje>=50% y <75%.
		Fuera de nivel:	Porcentaje<50%.
*/

package main
import "fmt"

func main(){
	var total, contestadas int
	var porcentaje float32

	fmt.Println("Total de preguntas del examen: ")
	fmt.Scan(&total)
	fmt.Println("Cantidad de preguntas respondidas: ")
	fmt.Scan(&contestadas)

	porcentaje = (float32(contestadas)/float32(total))*100

	if porcentaje >=float32(90){
		fmt.Println("Nivel Máximo.")
	}
	if porcentaje >=float32(75){
		if porcentaje  <  float32(90){
			fmt.Println("Nivel Medio.")
		}
	}
	if porcentaje >=float32(50){
		if porcentaje  <  float32(75){
			fmt.Println("Nivel Regular.")
		}
	}
	if porcentaje <float32(50){
		fmt.Println("Fuera De Nivel.")
	}
}
