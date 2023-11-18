package main

import (
	"fmt"
	"net"
)

const HOST string = "localhost"

var users []net.Conn


func main() {
	comms := make(chan string)
	listen, err := net.Listen("tcp", ":9999")

	
	if err != nil {
		fmt.Println("[!]Erro: ", err)
		return
	}
	
	fmt.Println("[+]Server running on", HOST, ":9999")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("[!]Erro: ", err)
			continue
		}
		users = append(users, conn)
		
		go handleConnection(conn)
		// go broadcastMessage(conn)
	}

}

func broadcastMessage(){
	for idx := range users{
		fmt.Println(users[idx])
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// go broadcastMessage(conn)
		buff := make([]byte, 1024)
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Println("[!]Erro: ", err)
			return
		}
		// broadcastMessage()
		buff <- conn
		fmt.Printf("[✅]Incoming: %s \n", buff);
		fmt.Printf("%s", users)
		conn.Write([]byte(buff))
	}
}