package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	stopCharacter = "2402"
)

func socketClient(ip string, port int, message string) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer conn.Close()

	conn.Write([]byte(message))
	conn.Write([]byte(stopCharacter))
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}

func main() {

	var (
		ip   = getIp()
		port = 8888
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	message, _ := reader.ReadString('\n')
	fmt.Println(message)

	socketClient(ip, port, message)

}
