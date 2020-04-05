package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	stopCharacter = "2402"
)

func socketServer(port int) {

	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))

	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}

	defer listen.Close()

	log.Printf("Begin listen port: %d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {

	defer conn.Close()

	var (
		buf     = make([]byte, 1024)
		r       = bufio.NewReader(conn)
		w       = bufio.NewWriter(conn)
		message string
	)
	n, err := r.Read(buf)
	data := string(buf[:n])

ILOOP:
	for {

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Receive:", data)
			//if isTransportOver(data) {
				data = data[:len(data)-1]
				message = fmt.Sprintf("Hello %s, welcome!", data)
				break ILOOP
			//}

		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}

	}
	w.Write([]byte(message))
	w.Flush()
	log.Printf("%s has spoken", data)

}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, stopCharacter)
	return
}

func main() {

	port := 8888

	socketServer(port)

}
