/*
Practica 03: Calcular el Precio de Venta

Enunciado: dado el valor de venta de productos, hallar el IGV (18%) y el precio de venta.

Análisis: para la solución de este problema, se requiere que el usuario ingrese el valor de venta del producto y el sistema realice el cálculo respectivo para hallar el IGV y el precio de venta.

Valor de venta: Valor introducido por el usuario.

IGV: Impuesto General a las Ventas o IGV y su tasa es establecida por el estado peruano hoy es del 18 %

Precio Venta:
El precio de venta es lo que debe pagar el cliente entonces significa que incluye el IGV, por lo tanto, la parte del dinero que le pertenece a la empresa se llama valor de venta.

Costo + Ganancias = Valor de Venta
Valor de Venta + IGV = Precio de Venta

Cuando un negocio compra se vuelve cliente de otro negocio, y se aplica la misma lógica, el negocio pagará al estado un IGV mediante su proveedor. El proveedor recibe el dinero y paga a nombre del negocio al estado, es decir, el precio que paga el negocio a su proveedor incluye el IGV, por lo tanto, la parte de dinero que pertenece al proveedor se le llama valor de compra.

Valor de Compra + IGV = Precio de Compra

... Sigan esta regla para construir su precio de venta: el IGV se adiciona al valor de venta y se excluye del valor de compra. Si no aplican esta regla cuando vendan, sacrificarán parte de sus ganancias por devolver dinero al estado, como si asumieran el IGV de sus clientes y ésa es la razón del porqué muchos piensan que el estado les quita mucho efectivo.

Extraido desde URI: https://www.misfinanzas.pe/igv-y-precios/

además, ¿Cómo se calcula el IGV?

El cálculo del IGV se hace aplicando la tasa del 18% sobre el valor de venta (también llamado base imponible).

    Si tenemos el Valor de Venta y queremos aplicar el 18%: IGV = VALOR DE VENTA x 0.18
    Si tenemos el Precio de Venta y queremos desglosar el 18%: IGV = PRECIO DE VENTA / 1.18

VALOR DE VENTA 	1,000.00
IGV 18%	180.00
PRECIO DE VENTA	1,180.00
*/

package main

// `fmt` es un paquete que contiene funciones para formatear e imprimir cadenas.
import "fmt"

func main() {
	//Definicion de variables: var nombreVariable tipo

    var valorVenta float64
    var precioVenta float64

		//Imprimir por pantalla y leer desde teclado
    fmt.Print("Ingrese el valor de venta del producto: ")
    fmt.Scan(&valorVenta)

   	precioVenta=valorVenta+(valorVenta*(0.18))

    fmt.Print("El Precio de venta debe ser de $ ",int64(precioVenta), " /moneda local.")
}


