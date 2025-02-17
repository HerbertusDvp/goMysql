package handlers

import (
	"fmt"
	"goMysql/conect"
	"goMysql/modelos"
	"log"
)

var pln = fmt.Println

func Listar() {
	conect.Conecta()
	query := "select*from cliente"

	datos, err := conect.Conexion.Query(query)

	if err != nil {
		pln(err)
	}

	defer conect.CerrarConexion()

	clientes := modelos.Clientes{}

	for datos.Next() {
		dato := modelos.Cliente{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		clientes = append(clientes, dato)
	}

	for _, valor := range clientes {
		pln(valor)
	}
}

func ListaById(id int) {

	conect.Conecta()
	query := "select * from cliente where id = ?"
	datos, err := conect.Conexion.Query(query, id)

	if err != nil {
		pln(err)
	}
	defer conect.CerrarConexion()

	clientes := modelos.Clientes{}

	for datos.Next() {
		var dato modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		clientes = append(clientes, dato)

		if err != nil {
			log.Fatal(err)
		}

		for _, valor := range clientes {
			pln(valor)
		}
	}

}
