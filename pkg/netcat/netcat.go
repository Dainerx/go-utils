package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	go func() {
		// copy ouput
		io.Copy(os.Stdout, conn) // copy server output
		log.Println("done")
		done <- true
	}()
	// copy input
	mustCopy(conn, os.Stdin)
	if tcpconn, ok := conn.(*net.TCPConn); ok {
		tcpconn.CloseWrite()
	}
	<-done
}
