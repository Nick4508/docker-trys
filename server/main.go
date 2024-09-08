package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("Servidor TCP escuchando en el puerto 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Cliente conectado:", conn.RemoteAddr().String())
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer mensaje:", err)
		return
	}
	fmt.Println("Mensaje recibido:", message)

	conn.Write([]byte("Mensaje recibido\n"))
}
