// Chat is a server that lets clients chat with each other.
// For n clients: 2n+2 concurrent communications
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	idleTimeout = time.Minute * 1
)

// Client is represented by his outgoing message channel
type Client struct {
	name         string
	channel      chan string
	lastActivity time.Time
}

var (
	entering = make(chan Client) // entering clients channel
	leaving  = make(chan Client) // leaving clients channel
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	// Set of outgoing clients channels (set of connected clients)
	clients := make(map[string]Client) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast every message to all connected clients
			for _, cli := range clients {
				cli.channel <- msg
			}
		case cli := <-entering:
			clients[cli.name] = cli // Add to the set this client

			nowOnline := "===========================\nOnline clients:\n"
			for name := range clients {
				nowOnline += name + "\t"
			}
			nowOnline += "\n===========================\n"

			for _, cli := range clients {
				cli.channel <- nowOnline
			}
		case cli := <-leaving:
			delete(clients, cli.name) // Remove client from set
			close(cli.channel)        // Close his outgoing channel
		}
	}
}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	var who string
	fmt.Fprint(conn, "Input your name: ")
	if input.Scan() {
		who = input.Text()
	}

	newClient := Client{
		name:         who,
		channel:      make(chan string),
		lastActivity: time.Now(),
	}

	go clientWriter(conn, newClient.channel)

	newClient.channel <- "You are " + newClient.name
	messages <- newClient.name + " has arrived"
	entering <- newClient

	// Go routine to detect idle clients
	go func() {
		for {
			d := time.Now().Sub(newClient.lastActivity)
			if d > idleTimeout {
				conn.Close() // ignore error
				return
			}
		}
	}()

	// Read input from the new client and broadcast it.
	for input.Scan() {
		messages <- newClient.name + ": " + input.Text()
		newClient.lastActivity = time.Now()
	}

	leaving <- newClient
	messages <- newClient.name + " has left"
	conn.Close()
}

// Write to client messages
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	// main routine only listens to new connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// spawn a new go routine to handle the client
		go handleConn(conn)
	}
}
