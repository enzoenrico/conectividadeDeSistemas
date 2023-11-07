package main

import (
	"fmt"
	"net"
)

const HOST string = "localhost"

func main() {
	listen, err := net.Listen("tcp", ":9999")

	fmt.Println("[+]Server running on", HOST, ":9999")

	if err != nil {
		fmt.Println("[!]Erro: ", err)
		return
	}

	// While loop do Go (linguagem meio estranha mas ta valendo)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("[!]Erro: ", err)
			continue
		}
		// Threading (retiro o que eu disse eu amo Go)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Fecha a conexão quando a função retornar algo
	defer conn.Close()
	for {
		buff := make([]byte, 1024)
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Println("[!]Erro: ", err)
			return
		}
		fmt.Printf("[✅]Incoming: %s \n", buff)
		// Da um encode no texto e joga pra uma array de bits
		// conn.Write([]byte("[SERVER] Recived"))
		conn.Write([]byte("test"))
	}
}
