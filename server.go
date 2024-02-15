package main

import (
	"io"
	"log"
	"net"
	"time"
)

func receiveTCPConn(ln *net.TCPListener) {
	for {
		err := ln.SetDeadline(time.Now().Add(time.Second * 100))
		if err != nil {
			log.Fatal(err)
		}
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go echoHandler(conn)
	}
}

func echoHandler(conn *net.TCPConn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, "Socket Connection!!\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal("1: ", err)
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("2: ", err)
	}
	receiveTCPConn(ln)

}
