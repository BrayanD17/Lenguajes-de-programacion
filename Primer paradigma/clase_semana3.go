package main

import (
	"fmt"
	"sort"
	"strings"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 20 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var existe bool = false
	for i := 0; i < len(*l); i++ {
		if strings.Compare(nombre, (*l)[i].nombre) == 0 {
			(*l)[i].cantidad = (*l)[i].cantidad + cantidad
			existe = true
		}
	}
	if existe == false {
		(*l) = append((*l), producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant
		}
		if (*l)[prod].cantidad == 0 {
			var i int = 0
			for i = 0; i < len(*l); i++ {
				if (*l)[i].nombre == nombre {
					(*l) = append((*l)[:i], (*l)[i+1:]...)
					fmt.Println("Producto eliminado!!")
				}
			}
		}

		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	}
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	var i int = 0
	var lminimos listaProductos
	for i = 0; i < len(*l); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			lminimos = append(lminimos, (*l)[i])
		}
	}
	// debe retornar una nueva lista con productos con existencia mínima
	return lminimos
}

func aumentarInventarioDeMinimos(l listaProductos) listaProductos {
	var i int = 0
	var Alminimos listaProductos
	for i = 0; i < len(l); i++ {
		Alminimos = append(Alminimos, producto{nombre: (l)[i].nombre, cantidad: existenciaMinima - (l)[i].cantidad, precio: (l)[i].precio})
	}
	return Alminimos
}

/*
func ordenamiento () listaProductos{
	var listaOrdenada listaProductos
	sort.SliceStable(listaProductos, func(i, j int) bool {return listaProductos[i].cantidad < listaProductos[j].cantidad})
}*/

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	lProductos.venderProducto("arroz", 4)
	fmt.Println(lProductos)
	lProductos.agregarProducto("azucar", 20, 1500)
	fmt.Println(lProductos)
	lProductos.venderProducto("frijoles", 4)
	fmt.Println(lProductos)
	lProductos.venderProducto("leche", 10)
	lProductos.agregarProducto("azucar", 20, 1500)
	fmt.Println(lProductos)
	fmt.Println("Lista con minimos")
	fmt.Println(lProductos.listarProductosMínimos())

	/* Ejercicios para repositorio*/

	fmt.Println("Productos a los cuales se les amplio el inventario en: ")
	var lminimos listaProductos = lProductos.listarProductosMínimos()
	fmt.Println(aumentarInventarioDeMinimos(lminimos))

	sort.SliceStable(lProductos, func(i, j int) bool { return lProductos[i].cantidad < lProductos[j].cantidad })
	fmt.Println("Lista ordenada por cantidad disponible")
	fmt.Println(lProductos)

}

// si ha sobrado tiempo antes del receso, el ejercicio se podría ampliar para que los los productos se almacenen en archivos de texto
// que al inicio se carguen del archivo a la lista y que al final se actualice el archivo con el contenido de la lista
