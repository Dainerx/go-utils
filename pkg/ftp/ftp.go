package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type file struct {
	name    string
	content []byte
}

type server struct {
	files []file
}

var ftp *server

func init() {
	f1 := file{"hello.txt", []byte("helllo world")}
	f2 := file{"goodbye.txt", []byte("goodbye")}

	ftp = &server{[]file{f1, f2}}
}

func (s *server) ls(c net.Conn) (int, error) {
	dir := ""
	for _, file := range s.files {
		dir += file.name + "\n"
	}
	return io.WriteString(c, dir)
}

func handle(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		command := scanner.Text()
		switch command {
		case "ls":
			_, err := ftp.ls(c)
			if err != nil {
				return
			}
		case "close":
			io.WriteString(c, "bye!")
			return
		default:
			_, err := io.WriteString(c, fmt.Sprintf("%s: %s\n", command, "command not found"))
			if err != nil {
				return
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handle(conn) // handle connections | |
	}

}
