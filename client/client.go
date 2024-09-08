package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Cambia "server-ip" por la IP de la máquina del servidor
	conn, err := net.Dial("tcp", "10.35.168.79:8080")
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Conectado al servidor, enviando mensaje...")
	fmt.Fprintf(conn, "Hola desde el cliente\n")

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Respuesta del servidor:", message)
}
