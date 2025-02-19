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
	pln("6. Salir")

	if scanner.Scan() {

		if scanner.Text() == "1" {
			Listar()

		} else if scanner.Text() == "2" {
			p("Ingresa el id del cliente: ")
			if scanner.Scan() {
				id, _ = strconv.Atoi(scanner.Text())
			}
			ListaById(id)

		} else if scanner.Text() == "3" {
			p("Ingresa un nombre: ")
			if scanner.Scan() {
				nombre = scanner.Text()
			}
			p("Ingresa el correo: ")
			if scanner.Scan() {
				correo = scanner.Text()
			}
			p("Ingresa el teléfono: ")
			if scanner.Scan() {
				telefono = scanner.Text()
			}

			cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
			Insertar(cliente)

		} else if scanner.Text() == "4" {
			p("Ingresa el id:")
			if scanner.Scan() {
				id, _ = strconv.Atoi(scanner.Text())
			}
			p("Ingresa el nombre: ")
			if scanner.Scan() {
				nombre = scanner.Text()
			}
			p("Ingresa el correo: ")
			if scanner.Scan() {
				correo = scanner.Text()
			}
			p("Ingresa el telefono:")
			if scanner.Scan() {
				telefono = scanner.Text()
			}
			cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
			Actualizar(cliente, id)

		} else if scanner.Text() == "5" {
			p("Ingresa el id del cliente a leiminar: ")
			if scanner.Scan() {
				id, _ = strconv.Atoi(scanner.Text())
			}
			Eliminar(id)
		} else if scanner.Text() == "6" {
			pln("Saliendo ...")
			return
		}
		Menu()
	}
}
