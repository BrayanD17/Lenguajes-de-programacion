package main

import (
	"fmt"
	"strings"
)

type calzado struct {
	modelo    string
	precio    int
	talla     int
	n_zapatos int
}

var lista_zapatos_array [10]calzado
var newlista_zapatos_array []calzado

func agregarInventarioQ() {
	lista_zapatos_array[0].modelo = "Christian Pepe"
	lista_zapatos_array[0].precio = 30000
	lista_zapatos_array[0].talla = 38
	lista_zapatos_array[0].n_zapatos = 12

	lista_zapatos_array[1].modelo = "Jimmy"
	lista_zapatos_array[1].precio = 60000
	lista_zapatos_array[1].talla = 40
	lista_zapatos_array[1].n_zapatos = 1

	lista_zapatos_array[2].modelo = "Aquazzura"
	lista_zapatos_array[2].precio = 120000
	lista_zapatos_array[2].talla = 42
	lista_zapatos_array[2].n_zapatos = 10

	lista_zapatos_array[3].modelo = "Gianvito"
	lista_zapatos_array[3].precio = 50000
	lista_zapatos_array[3].talla = 36
	lista_zapatos_array[3].n_zapatos = 20

	lista_zapatos_array[4].modelo = "Stuart"
	lista_zapatos_array[4].precio = 20000
	lista_zapatos_array[4].talla = 44
	lista_zapatos_array[4].n_zapatos = 4

	lista_zapatos_array[5].modelo = "Charlotte"
	lista_zapatos_array[5].precio = 70000
	lista_zapatos_array[5].talla = 34
	lista_zapatos_array[5].n_zapatos = 13

	lista_zapatos_array[6].modelo = "Roger"
	lista_zapatos_array[6].precio = 40000
	lista_zapatos_array[6].talla = 35
	lista_zapatos_array[6].n_zapatos = 24

	lista_zapatos_array[7].modelo = "Alexandre"
	lista_zapatos_array[7].precio = 65000
	lista_zapatos_array[7].talla = 37
	lista_zapatos_array[7].n_zapatos = 2

	lista_zapatos_array[8].modelo = "Giuseppe"
	lista_zapatos_array[8].precio = 90000
	lista_zapatos_array[8].talla = 39
	lista_zapatos_array[8].n_zapatos = 5

	lista_zapatos_array[9].modelo = "Salvatore"
	lista_zapatos_array[9].precio = 110000
	lista_zapatos_array[9].talla = 40
	lista_zapatos_array[9].n_zapatos = 10
}

func venta() {
	for i := 0; i < len(newlista_zapatos_array); i++ {
		fmt.Println("Modelo: ", newlista_zapatos_array[i].modelo)
		fmt.Println("Cantidad disponible: ", newlista_zapatos_array[i].n_zapatos)
		fmt.Println("------------------------------")
	}

	fmt.Print("Escriba el nombre del zapato que desea comprar: ")
	var modelo string
	fmt.Scanf("%(^\n)", &modelo)
	fmt.Print("Zapatos disponibles!!!!")
	for i := 0; i < len(newlista_zapatos_array); i++ {
		if strings.Compare(modelo, newlista_zapatos_array[i].modelo) == 0 /*modelo == lista_zapatos_array[i].modelo*/ {
			lista_zapatos_array[i].n_zapatos = newlista_zapatos_array[i].n_zapatos - 1
			fmt.Println("Compra realizada!!!")
			fmt.Println("Modelo: ", newlista_zapatos_array[i].modelo)
			fmt.Println("Cantidad disponible: ", newlista_zapatos_array[i].n_zapatos)
			fmt.Println("----------")
		}
		if newlista_zapatos_array[i].n_zapatos == 0 {
			newlista_zapatos_array = eliminar(i)
			fmt.Println("Zapato agotado!!")
		}
	}
}
func eliminar(indice int) []calzado {
	return append(newlista_zapatos_array[:indice], newlista_zapatos_array[indice+1:]...)
}
func main() {
	agregarInventarioQ()
	newlista_zapatos_array = lista_zapatos_array[:]
	fmt.Println("Zapatos disponibles!!!!")
	venta()
	for true {
		fmt.Println("Desea volver a comprar:\n1) si\n2) no\n")
		var opcion string
		fmt.Scanln(&opcion)
		if opcion == "si" {
			venta()
		}
		if opcion == "no" {
			return
		}
	}
}
