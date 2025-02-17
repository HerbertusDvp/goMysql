package main

import (
	"fmt"
	"goMysql/handlers"
)

var pln = fmt.Println

func main() {
	//Todos los clientes
	pln("Lista de todos los clientes: \n")

	handlers.Listar()

	pln("\nListado de clientes por ID: \n")

	handlers.ListaById(2)

}
