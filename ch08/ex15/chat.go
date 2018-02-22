package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	who string
	ch  chan<- string // an outgoing message channel
}

const limitTimeout = 5 * time.Minute

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.ch <- msg:
				default:
					// if cli.ch cannot receive msg, skip cli
				}
			}

		case cli := <-entering:
			if len(clients) > 0 {
				cli.ch <- "active users: "
				for other := range clients {
					cli.ch <- "\t" + other.who
				}
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 30) // outgoing client messages: buffer 30
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)

	who := conn.RemoteAddr().String()

	ch <- "Input your name"
	if input.Scan() {
		name := input.Text()
		if len(name) > 0 {
			who = name
		}
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{who, ch}

	timer := time.AfterFunc(limitTimeout, func() {
		conn.Close()
	})

	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(limitTimeout)
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client{who, ch}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
