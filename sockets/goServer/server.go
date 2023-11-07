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

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("[!]Erro: ", err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buff := make([]byte, 1024)
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Println("[!]Erro: ", err)
			return
		}
		fmt.Printf("[✅]Incoming: %s \n", buff);
		conn.Write([]byte("[127.0.0.1] OK"))
	}
}
