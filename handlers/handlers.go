package handlers

import (
	"bufio"
	"fmt"
	"goMysql/conect"
	"goMysql/modelos"
	"log"
	"os"
	"strconv"
)

var pln = fmt.Println
var p = fmt.Print

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

	resultado := true
	conect.Conecta()
	query := "select * from cliente where id = ?"
	datos, err := conect.Conexion.Query(query, id)

	if err != nil {
		pln(err)
	}
	defer conect.CerrarConexion()

	clientes := modelos.Clientes{}

	for datos.Next() {
		resultado = false
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

	if resultado {
		pln("No se encontraron coinsidencias..")
	}
}

func Insertar(cliente modelos.Cliente) {
	conect.Conecta()
	query := "insert into cliente values (null, ?, ?, ?)"
	_, err := conect.Conexion.Exec(query, cliente.Nombre, cliente.Correo, cliente.Telefono)

	if err != nil {
		panic(err)
	}
	//pln(result)
	pln("Se registró el cliente: ", cliente)
}

func Actualizar(cliente modelos.Cliente, id int) {
	conect.Conecta()
	query := "update cliente set nombre = ?, correo=?, telefono=? where id = ?"
	_, err := conect.Conexion.Exec(query, cliente.Nombre, cliente.Correo, cliente.Telefono, id)

	if err != nil {
		panic(err)
	}
	//pln(result)
	pln("Registro modificado")
}

func Eliminar(id int) {
	conect.Conecta()
	query := "delete from cliente where id = ?"
	_, err := conect.Conexion.Exec(query, id)

	if err != nil {
		panic(err)
	}
	pln("Se elminó el cliente con id: ", id)
}

var id int
var nombre, correo, telefono string

func Menu() {
	scanner := bufio.NewScanner(os.Stdin)

	pln("-- Menú de opiones --\n")
	pln("1. Listar clientes")
	pln("2. Buscar cliente por id")
	pln("3. Registrar cliente")
	pln("4. Modificar cliente")
	pln("5. Eliminar cliente")

	if scanner.Scan() {
		for {
			if scanner.Text() == "1" {
				Listar()
				return
			}
			if scanner.Text() == "2" {
				p("Ingresa el id del cliente: ")
				if scanner.Scan() {
					id, _ = strconv.Atoi(scanner.Text())
				}
				ListaById(id)
				return
			}
			if scanner.Text() == "3" {

			}
		}
	}
}
